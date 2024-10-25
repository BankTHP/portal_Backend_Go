package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, userHandler *handler.UserHandlers, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/users")
	api.Post("/create", authMiddleware.HasRole("client_user", "client_admin"), userHandler.CreateUser)
	api.Put("/update", authMiddleware.HasRole("client_user", "client_admin"), userHandler.UpdateUserInfo)
	api.Get("/info/:userId", authMiddleware.HasRole("client_user", "client_admin"), userHandler.GetUserInfoByUserId)
	api.Post("/check", authMiddleware.HasRole("client_user", "client_admin"), userHandler.CheckUser)
}
