package handler

import (
	"fmt"
	"path/filepath"
	"strings"
	"pccth/portal-blog/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type VideoHandler struct {
	uploadPath string
}

func NewVideoHandler(uploadPath string) *VideoHandler {
	return &VideoHandler{
		uploadPath: uploadPath,
	}
}

func (h *VideoHandler) UploadVideo(c *fiber.Ctx) error {
	// รับไฟล์จาก request
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่พบไฟล์วิดีโอ กรุณาเลือกไฟล์วิดีโอที่ต้องการอัปโหลด",
		})
	}

	// ตรวจสอบขนาดไฟล์
	maxSize := viper.GetInt64("app.upload.max_size")
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   fmt.Sprintf("ไม่สามารถอัปโหลดได้: ขนาดไฟล์ %.2f MB เกินขนาดที่กำหนด (ไม่เกิน %.2f MB)", 
				float64(file.Size)/(1024*1024), 
				float64(maxSize)/(1024*1024)),
		})
	}

	// ตรวจสอบนามสกุลไฟล์
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedTypes := viper.GetStringSlice("app.upload.allowed_types")
	isAllowed := false
	for _, allowedType := range allowedTypes {
		if ext == allowedType {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   fmt.Sprintf("นามสกุลไฟล์ไม่ถูกต้อง รองรับเฉพาะ: %s", strings.Join(allowedTypes, ", ")),
		})
	}

	// สร้างชื่อไฟล์ใหม่
	filename := uuid.New().String() + ext
	fullPath := filepath.Join(h.uploadPath, filename)

	// บันทึกไฟล์
	if err := c.SaveFile(file, fullPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   "เกิดข้อผิดพลาดในการอัปโหลดไฟล์: " + err.Error(),
		})
	}

	// สร้าง URL สำหรับเข้าถึงวิดีโอ
	baseURL := fmt.Sprintf("http://localhost:%s", viper.GetString("app.port"))
	videoURL := fmt.Sprintf("%s/videos/%s", baseURL, filename)

	return c.JSON(model.VideoResponse{
		Success: true,
		FullURL: videoURL,
		Message: fmt.Sprintf("อัปโหลดวิดีโอสำเร็จ (ขนาด: %.2f MB)", float64(file.Size)/(1024*1024)),
	})
} 