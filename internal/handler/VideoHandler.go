package handler

import (
	"fmt"
	"path/filepath"
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
			Error:   "ไม่พบไฟล์วิดีโอ",
		})
	}

	// ตรวจสอบนามสกุลไฟล์
	ext := filepath.Ext(file.Filename)
	if ext != ".mp4" && ext != ".avi" && ext != ".mov" {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "รองรับเฉพาะไฟล์ .mp4, .avi, และ .mov เท่านั้น",
		})
	}

	// สร้างชื่อไฟล์ใหม่ด้วย UUID
	filename := uuid.New().String() + ext
	filepath := fmt.Sprintf("/videos/%s", filename)
	
	// บันทึกไฟล์
	err = c.SaveFile(file, fmt.Sprintf("./uploads%s", filepath))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่สามารถบันทึกไฟล์ได้",
		})
	}

	// สร้าง full URL โดยใช้ domain จาก config
	port := viper.GetString("app.port")
	fullURL := fmt.Sprintf("http://localhost:%s%s", port, filepath)

	return c.JSON(model.VideoResponse{
		Success: true,
		FullURL: fullURL,
	})
} 