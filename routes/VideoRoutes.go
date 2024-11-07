package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func VideoRoutes(app *fiber.App, videoHandler *handler.VideoHandler, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/upload")
	api.Post("/video", authMiddleware.HasRole("client_user", "client_admin"), videoHandler.UploadVideo)
} 