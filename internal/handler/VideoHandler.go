package handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
	// ตรวจสอบว่ามี ffprobe หรือไม่
	hasFfprobe := checkFfprobeExists()

	// รับไฟล์จาก request
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่พบไฟล์วิดีโอ",
		})
	}

	// ตรวจสอบขนาดไฟล์
	maxSize := viper.GetInt64("app.upload.max_size")
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   fmt.Sprintf("ขนาดไฟล์ต้องไม่เกิน %.2f MB", float64(maxSize)/(1024*1024)),
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
			Error:   fmt.Sprintf("รองรับเฉพาะไฟล์: %s", strings.Join(allowedTypes, ", ")),
		})
	}

	// สร้างชื่อไฟล์ใหม่
	filename := uuid.New().String() + ext
	fullPath := filepath.Join(h.uploadPath, filename)

	// บันทึกไฟล์
	if err := c.SaveFile(file, fullPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่สามารถบันทึกไฟล์ได้: " + err.Error(),
		})
	}

	// ตรวจสอบความยาวของวิดีโอ ถ้ามี ffprobe
	if hasFfprobe {
		duration, err := getVideoDuration(fullPath)
		if err != nil {
			os.Remove(fullPath) // ลบไฟล์ถ้าเกิดข้อผิดพลาด
			return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
				Success: false,
				Error:   "ไม่สามารถตรวจสอบความยาวของวิดีโอได้: " + err.Error(),
			})
		}

		maxDuration := viper.GetFloat64("app.upload.max_duration")
		if duration > maxDuration {
			os.Remove(fullPath)
			return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
				Success: false,
				Error:   fmt.Sprintf("ความยาวของวิดีโอต้องไม่เกิน %.0f นาที", maxDuration/60),
			})
		}
	}

	// สร้าง URL สำหรับเข้าถึงวิดีโอ
	baseURL := fmt.Sprintf("http://localhost:%s", viper.GetString("app.port"))
	videoURL := fmt.Sprintf("%s/videos/%s", baseURL, filename)

	return c.JSON(model.VideoResponse{
		Success: true,
		FullURL: videoURL,
	})
}

// ฟังก์ชันตรวจสอบว่ามี ffprobe หรือไม่
func checkFfprobeExists() bool {
	_, err := exec.LookPath("ffprobe")
	return err == nil
}

func getVideoDuration(filepath string) (float64, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filepath)

	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("ffprobe error: %v", err)
	}

	duration, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		return 0, fmt.Errorf("invalid duration format: %v", err)
	}

	return duration, nil
} 