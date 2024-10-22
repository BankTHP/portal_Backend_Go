package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(app *fiber.App, commentHandler *handler.CommentHandlers) {

	api := app.Group("/comment")
	api.Post("/createComment", commentHandler.CreateComment)
	api.Get("/getCommentById/:id", commentHandler.GetCommentByID)
	api.Get("/getCommentByPostId/:id", commentHandler.GetCommentByPostID)
	api.Delete("/deleteComment/:id", commentHandler.DeleteComment)
	api.Get("/getPaginatedCommentsByPostId", commentHandler.GetPaginatedComments)
	api.Get("/getCommentByUserId", commentHandler.GetPaginatedCommentsByUserId)
	api.Put("/updateComment/:id", commentHandler.UpdateComment)

}
