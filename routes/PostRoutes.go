package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoutes(app *fiber.App, postHandler *handler.PostHandlers) {
    app.Post("/posts", postHandler.CreatePost)            
    app.Get("/posts/:id", postHandler.GetPostByID)       
    app.Put("/posts/:id", postHandler.UpdatePost)         
    app.Delete("/posts/:id", postHandler.DeletePost)      
    app.Get("/posts", postHandler.GetAllPosts)           
}

