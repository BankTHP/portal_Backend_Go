package handler

import (
	"pccth/portal-blog/internal/model"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService *service.UserService
}

func NewUserHandlers(userService *service.UserService) *UserHandlers {
	return &UserHandlers{userService: userService}
}

func (h *UserHandlers) CreateUser(ctx *fiber.Ctx) error {
	var createUserRequest model.CreateUserRequest
	if err := ctx.BodyParser(&createUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := h.userService.CreateUser(&createUserRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CREATE_USER_ERROR", err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
		"message": "สร้างผู้ใช้สำเร็จ",
	}))
}

func (h *UserHandlers) UpdateUserInfo(ctx *fiber.Ctx) error {
	var updateUserRequest model.UpdateUserRequest
	if err := ctx.BodyParser(&updateUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	if err := h.userService.UpdateUserInfo(&updateUserRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("UPDATE_USER_ERROR", err.Error()))
	}

	return ctx.JSON(model.NewSuccessResponse(map[string]string{
		"message": "อัปเดตข้อมูลผู้ใช้สำเร็จ",
	}))
}

func (h *UserHandlers) GetUserInfoByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	user, err := h.userService.GetUserInfoByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(model.NewErrorResponse("USER_NOT_FOUND", "ไม่พบผู้ใช้"))
	}

	return ctx.JSON(model.NewSuccessResponse(user))
}

func (h *UserHandlers) CheckUser(ctx *fiber.Ctx) error {
	var userInfo model.CreateUserRequest
	if err := ctx.BodyParser(&userInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.NewErrorResponse("INVALID_INPUT", "ข้อมูลไม่ถูกต้อง"))
	}

	exists, err := h.userService.CheckUser(&userInfo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.NewErrorResponse("CHECK_USER_ERROR", err.Error()))
	}

	if exists {
		return ctx.Status(fiber.StatusOK).JSON(model.NewSuccessResponse(map[string]string{
			"message": "ผู้ใช้มีอยู่แล้ว",
		}))
	} else {
		return ctx.Status(fiber.StatusCreated).JSON(model.NewSuccessResponse(map[string]string{
			"message": "สร้างผู้ใช้สำเร็จ",
		}))
	}
}
