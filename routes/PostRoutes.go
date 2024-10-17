package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, postHandler *handler.PostHandlers) {

	api := app.Group("/posts")
	api.Post("/createPost", postHandler.CreatePost)
	api.Get("/getPostById/:id", postHandler.GetPostByID)
	api.Put("/updatePost/:id", postHandler.UpdatePost)
	api.Delete("/deletePost/:id", postHandler.DeletePost)
	api.Get("/getAllPosts", postHandler.GetAllPosts)
	api.Get("/getPaginatedPosts", postHandler.GetPaginatedPosts)

}
