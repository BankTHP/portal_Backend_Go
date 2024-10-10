package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
    postService := service.NewPostService(db)
    postController := handler.NewPostHandlers(postService)
    PostRoutes(app, postController)
}
