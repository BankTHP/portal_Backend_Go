package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(app *fiber.App, commentHandler *handler.CommentHandlers, authMiddleware *middleware.AuthMiddleware) {

	api := app.Group("/comment")
	api.Post("/createComment", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.CreateComment)
	api.Get("/getCommentById/:id", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.GetCommentByID)
	api.Get("/getCommentByPostId/:id", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.GetCommentByPostID)
	api.Delete("/deleteComment/:id", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.DeleteComment)
	api.Get("/getPaginatedCommentsByPostId", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.GetPaginatedComments)
	api.Get("/getCommentByUserId", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.GetPaginatedCommentsByUserId)
	api.Put("/updateComment/:id", authMiddleware.HasRole("client_user", "client_admin"), commentHandler.UpdateComment)

}
