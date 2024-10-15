package handler

import (
	"pccth/portal-blog/internal/service"
	"pccth/portal-blog/internal/model"

	"github.com/gofiber/fiber/v2"
)

type ReleaseNoteHandler struct {
	releaseNoteService *service.ReleaseNoteService
}

func NewReleaseNoteHandler(releaseNoteService *service.ReleaseNoteService) *ReleaseNoteHandler {
	return &ReleaseNoteHandler{releaseNoteService: releaseNoteService }
}


func (c *ReleaseNoteHandler) CreateReleaseNote(ctx *fiber.Ctx) error {
	var createReleaseNoteRequest model.CreateReleaseNoteRequest
	if err := ctx.BodyParser(&createReleaseNoteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid input"})
	}

	if err := c.releaseNoteService.CreateRelease(&createReleaseNoteRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(map[string]interface{}{"message": "Release note created successfully"})
}
