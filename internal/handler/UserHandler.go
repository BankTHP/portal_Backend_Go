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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.userService.CreateUser(&createUserRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandlers) UpdateUserInfo(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	var updateUserRequest model.UpdateUserRequest
	if err := ctx.BodyParser(&updateUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := h.userService.UpdateUserInfo(userId, &updateUserRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "User updated successfully"})
}

func (h *UserHandlers) GetUserInfoByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")

	user, err := h.userService.GetUserInfoByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return ctx.JSON(user)
}



func (h *UserHandlers) CheckUser(ctx *fiber.Ctx) error {
	var userInfo model.CreateUserRequest
	if err := ctx.BodyParser(&userInfo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}

	exists, err := h.userService.CheckUser(&userInfo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if exists {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "ผู้ใช้มีอยู่แล้ว"})
	} else {
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "สร้างผู้ใช้สำเร็จ"})
	}
}
