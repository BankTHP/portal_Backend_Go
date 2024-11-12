package routes

import (
	"log"
	"os"
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func VideoRoutes(app *fiber.App, videoHandler *handler.VideoHandler, authMiddleware *middleware.AuthMiddleware) {

	uploadDir := "./uploads/videos"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("ไม่สามารถสร้าง upload directory ได้: %v", err)
	}

	app.Static("/videos", "./uploads/videos")

	api := app.Group("/video")
	api.Post("/upload",authMiddleware.HasRole("client_user", "client_admin"), videoHandler.UploadVideo)
	api.Get("/getvdoByName/:name",authMiddleware.HasRole("client_user", "client_admin"), videoHandler.GetVideoByName)

}
