package handler

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"pccth/portal-blog/internal/model"
	"strings"

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

func ensureDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.MkdirAll(dirPath, 0755)
	}
	return nil
}

func getVideoDuration(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", fmt.Errorf("ไม่พบไฟล์: %v", err)
	}

	cmd := exec.Command("ffmpeg", "-i", filename, "2>&1")
	output, err := cmd.CombinedOutput()
	if err != nil {
		outputStr := string(output)
		for _, line := range strings.Split(outputStr, "\n") {
			if strings.Contains(line, "Duration:") {
				parts := strings.Split(line, "Duration: ")
				if len(parts) > 1 {
					duration := strings.Split(parts[1], ",")[0]

					return strings.TrimSpace(duration), nil
				}
			}
		}
	}

	return "", fmt.Errorf("ไม่สามารถอ่านความยาววิดีโอได้")
}

func (h *VideoHandler) UploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่พบไฟล์วิดีโอ กรุณาเลือกไฟล์วิดีโอที่ต้องการอัปโหลด",
		})
	}

	maxSize := int64(500 * 1024 * 1024) // 500MB
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error: fmt.Sprintf("ไม่สามารถอัปโหลดได้: ขนาดไฟล์ %.2f MB เกินขนาดที่กำหนด (ไม่เกิน %.2f MB)",
				float64(file.Size)/(1024*1024),
				float64(maxSize)/(1024*1024)),
		})
	}

	// ตรวจสอบนามสกุลไฟล์แบบง่าย
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".mp4" && ext != ".mov" && ext != ".avi" && ext != ".wmv" {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "นามสกุลไฟล์ไม่ถูกต้อง รองรับเฉพาะ: .mp4, .mov, .avi, .wmv",
		})
	}

	uploadDir := filepath.Join(".", "uploads", "videos")
	if err := ensureDir(uploadDir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่สามารถสร้างโฟลเดอร์ได้",
		})
	}

	filename := uuid.New().String() + ext
	urlPath := filepath.Join("/videos", filename)
	fullPath := filepath.Join(h.uploadPath, filename)

	// Save file
	if err := c.SaveFile(file, fullPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   "เกิดข้อผิดพลาดในการอัปโหลดไฟล์: " + err.Error(),
		})
	}

	duration, err := getVideoDuration(fullPath)
	if err != nil {
		fmt.Printf("ไม่สามารถอ่านความยาววิดีโอได้: %v\n", err)
		duration = "ไม่สามารถอ่านความยาววิดีโอได้"
	} else {
		parts := strings.Split(duration, ":")
		if len(parts) == 3 {
			hours := strings.TrimSpace(parts[0])
			if hours != "00" {
				os.Remove(fullPath)
				return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
					Success: false,
					Error:   "วิดีโอต้องมีความยาวไม่เกิน 1 ชั่วโมง",
				})
			}
		}
	}

	fileSizeMB := fmt.Sprintf("%.2f MB", float64(file.Size)/1024/1024)
	port := viper.GetString("app.port")
	fullURL := fmt.Sprintf("http://localhost:%s%s", port, urlPath)

	return c.JSON(model.VideoResponse{
		Success:  true,
		FullURL:  fullURL,
		Duration: duration,
		Size:     fileSizeMB,
	})
}
