package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func NewsRoutes(app *fiber.App, newsHandler *handler.NewsHandlers) {

    api := app.Group("/news")
    api.Post("/createNews", newsHandler.CreateNews)            
    api.Get("/getNewsById/:id", newsHandler.GetNewsByID)       
    api.Put("/updateNews/:id", newsHandler.UpdateNews)         
    api.Delete("/deleteNews/:id", newsHandler.DeleteNews)      
    api.Get("/getAllNews", newsHandler.GetAllNews)           
}

