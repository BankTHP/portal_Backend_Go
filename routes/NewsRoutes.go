package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func NewsRoutes(app *fiber.App, newsHandler *handler.NewsHandlers) {
    app.Post("/createNews", newsHandler.CreateNews)            
    app.Get("/getNewsById/:id", newsHandler.GetNewsByID)       
    app.Put("/updateNews/:id", newsHandler.UpdateNews)         
    app.Delete("/deleteNews/:id", newsHandler.DeleteNews)      
    app.Get("/getAllNews", newsHandler.GetAllNews)           
}

