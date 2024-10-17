package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ReleaseNoteHandler struct {
	releaseNoteService *service.ReleaseNoteService
}

func NewReleaseNoteHandler(releaseNoteService *service.ReleaseNoteService) *ReleaseNoteHandler {
	return &ReleaseNoteHandler{releaseNoteService: releaseNoteService}
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

func (c *ReleaseNoteHandler) GetReleaseByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	println(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	release, err := c.releaseNoteService.GetReleaseByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Release note not found"})
	}

	return ctx.JSON(release)
}

func (c *ReleaseNoteHandler) UpdateRelease(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var updateRequest model.UpdateReleaseNoteRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := c.releaseNoteService.UpdateRelease(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Release Note updated successfully"})
}

func (c *ReleaseNoteHandler) DeleteRelease(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"error": "Invalid ID"})
	}

	if err := c.releaseNoteService.DeleteRelease(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(map[string]interface{}{"message": "Release note deleted successfully"})
}

func (c *ReleaseNoteHandler) GetAllRelease(ctx *fiber.Ctx) error {
	releases, err := c.releaseNoteService.GetAllRelease()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{"error": err.Error()})
	}

	return ctx.JSON(releases)
}
