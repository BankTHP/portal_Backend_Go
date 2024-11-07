package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB, authMiddleware *middleware.AuthMiddleware) {
	postService := service.NewPostService(db)
	postHandler := handler.NewPostHandlers(postService)
	PostRoutes(app, postHandler, authMiddleware)

	releaseNoteService := service.NewReleaseNoteService(db)
	releaseNoteHandler := handler.NewReleaseNoteHandler(releaseNoteService)
	ReleaseNoteRoutes(app, releaseNoteHandler, authMiddleware)

	commentService := service.NewCommentService(db)
	commentController := handler.NewCommentHandlers(commentService)
	CommentRoutes(app, commentController, authMiddleware)

	newsService := service.NewsPService(db)
	newsHandler := handler.NewsPHandlers(newsService)
	NewsRoutes(app, newsHandler, authMiddleware)

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandlers(userService)
	UserRoutes(app, userHandler, authMiddleware)

	// เพิ่ม video routes
	videoHandler := handler.NewVideoHandler("./uploads/videos")
	VideoRoutes(app, videoHandler, authMiddleware)
}
