package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, postHandler *handler.PostHandlers) {
	app.Post("/createPost", postHandler.CreatePost)
	app.Get("/getPostById/:id", postHandler.GetPostByID)
	app.Put("/updatePost/:id", postHandler.UpdatePost)
	app.Delete("/deletePost/:id", postHandler.DeletePost)
	app.Get("/getAllPosts", postHandler.GetAllPosts)
	app.Get("/getAllPostsByPage", postHandler.GetPaginatedPosts)

}
