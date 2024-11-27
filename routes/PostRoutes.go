package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, postHandler *handler.PostHandlers, authMiddleware *middleware.AuthMiddleware) {

	api := app.Group("/posts")
	api.Post("/createPost",	authMiddleware.HasRole("client_user", "client_admin"),  postHandler.CreatePost)
	api.Get("/getPostById/:id",authMiddleware.HasRole("client_user", "client_admin"), postHandler.GetPostByID)
	api.Put("/updatePost/:id", authMiddleware.HasRole("client_user", "client_admin"), postHandler.UpdatePost)
	api.Delete("/deletePost/:id", authMiddleware.HasRole("client_user", "client_admin"), postHandler.DeletePost)
	api.Get("/getAllPosts", authMiddleware.HasRole("client_user", "client_admin"), postHandler.GetAllPosts)
	api.Get("/getPaginatedPosts", authMiddleware.HasRole("client_user", "client_admin"), postHandler.GetPaginatedPosts)
	api.Get("/getAllPostByUserId", authMiddleware.HasRole("client_user", "client_admin"), postHandler.GetPaginatedPostsByUserId)

}
