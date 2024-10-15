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

// CreateReleaseNote godoc
// @Summary สร้างบันทึกการเผยแพร่ใหม่
// @Description สร้างบันทึกการเผยแพร่ใหม่ในระบบ
// @Tags release-notes
// @Accept json
// @Produce json
// @Param request body model.CreateReleaseNoteRequest true "ข้อมูลบันทึกการเผยแพร่"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /release-notes [post]
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
