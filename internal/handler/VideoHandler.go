package handler

import (
	"fmt"
	"path/filepath"
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type VideoHandler struct {
	videoService *service.VideoService
}

func NewVideoHandler(videoService *service.VideoService) *VideoHandler {
	return &VideoHandler{
		videoService: videoService,
	}
}

func (h *VideoHandler) UploadVideo(c *fiber.Ctx) error {
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "ไม่พบไฟล์วิดีโอ กรุณาเลือกไฟล์วิดีโอที่ต้องการอัปโหลด",
		})
	}

	uploadedFile := &model.UploadedFile{
		File:     file,
		Filename: file.Filename,
		Size:     file.Size,
		SaveFunc: func(path string) error {
			return c.SaveFile(file, path)
		},
	}

	response, err := h.videoService.ProcessVideoUpload(uploadedFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.VideoResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	if !response.Success {
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	
	port := viper.GetString("app.port")
	host := viper.GetString("app.host")
	response.FullURL = "http://" + host + ":" + port + filepath.ToSlash(response.FullURL)

	return c.JSON(response)
}

func (h *VideoHandler) GetVideoByName(c *fiber.Ctx) error {
	vdoName := c.Params("name")
	if vdoName == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.VideoResponse{
			Success: false,
			Error:   "กรุณาระบุชื่อวิดีโอ",
		})
	}

	video, err := h.videoService.GetVideoByName(vdoName)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(model.VideoResponse{
			Success: false,
			Error:   err.Error(),
		})
	}

	port := viper.GetString("app.port")
	host := viper.GetString("app.host")
	fullURL := fmt.Sprintf("http://%s:%s/videos/%s", host,port, video.VdoName)

	return c.JSON(model.VideoResponse{
		Success:  true,
		FullURL:  fullURL,
		Duration: video.VdoDuration,
		Size:     video.VdoSize,
	})
}
