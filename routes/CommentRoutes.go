package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(app *fiber.App, commentHandler *handler.CommentHandlers) {
	app.Post("/createComment", commentHandler.CreateComment)
	app.Get("/getCommentById/:id", commentHandler.GetCommentByID)
	app.Get("/getCommentByPostId/:id", commentHandler.GetCommentByPostID)
	app.Delete("/deleteComment/:id", commentHandler.DeleteComment)
	app.Get("/getAllCommentByPage", commentHandler.GetPaginatedComments)
	app.Put("/updateComment/:id", commentHandler.UpdateComment)

}
