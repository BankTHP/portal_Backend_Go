package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *handler.UserHandlers) {
	api := app.Group("/users")
	api.Post("/create", userHandler.CreateUser)
	api.Put("/update/:userId", userHandler.UpdateUserInfo)
	api.Get("/info/:userId", userHandler.GetUserInfoByUserId)
	api.Post("/check", userHandler.CheckUser)
}
