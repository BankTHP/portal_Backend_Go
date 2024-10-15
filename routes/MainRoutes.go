package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	postService := service.NewPostService(db)
	postHandler := handler.NewPostHandlers(postService)
	PostRoutes(app, postHandler)

	releaseNoteService := service.NewReleaseNoteService(db)
	releaseNoteHandler := handler.NewReleaseNoteHandler(releaseNoteService)
	ReleaseNoteRoutes(app, releaseNoteHandler)

	commentService := service.NewCommentService(db)
	commentController := handler.NewCommentHandlers(commentService)
	CommentRoutes(app, commentController)

	newsService := service.NewsPService(db)
	newsHandler := handler.NewsPHandlers(newsService)
	NewsRoutes(app, newsHandler)
}
