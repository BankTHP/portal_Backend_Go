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
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := c.releaseNoteService.CreateRelease(&createReleaseNoteRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CREATE_RELEASE_NOTE_ERROR", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
		"message": "สร้าง Release Note สำเร็จ",
	}))
}

func (c *ReleaseNoteHandler) GetReleaseByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	release, err := c.releaseNoteService.GetReleaseByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.NewErrorResponse("RELEASE_NOT_FOUND", "ไม่พบ Release Note"))
	}

	return ctx.JSON(model.NewSuccessResponse(release))
}

func (c *ReleaseNoteHandler) UpdateRelease(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	var updateRequest model.UpdateReleaseNoteRequest
	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := c.releaseNoteService.UpdateRelease(uint(id), &updateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("UPDATE_RELEASE_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "อัปเดต Release Note สำเร็จ",
	}))
}

func (c *ReleaseNoteHandler) DeleteRelease(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_ID", "รหัสไม่ถูกต้อง"))
	}

	if err := c.releaseNoteService.DeleteRelease(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("DELETE_RELEASE_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "ลบ Release Note สำเร็จ",
	}))
}

func (c *ReleaseNoteHandler) GetAllRelease(ctx *fiber.Ctx) error {
	releases, err := c.releaseNoteService.GetAllRelease()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("FETCH_RELEASES_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(releases))
}
