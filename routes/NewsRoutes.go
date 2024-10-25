package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewsRoutes(app *fiber.App, newsHandler *handler.NewsHandlers, authMiddleware *middleware.AuthMiddleware) {

	api := app.Group("/news")
	api.Post("/createNews", authMiddleware.HasRole("client_user", "client_admin"), newsHandler.CreateNews)
	api.Get("/getNewsById/:id", authMiddleware.HasRole("client_user", "client_admin"), newsHandler.GetNewsByID)
	api.Put("/updateNews/:id", authMiddleware.HasRole("client_user", "client_admin"), newsHandler.UpdateNews)
	api.Delete("/deleteNews/:id", authMiddleware.HasRole("client_user", "client_admin"), newsHandler.DeleteNews)
	api.Get("/getAllNews", authMiddleware.HasRole("client_user", "client_admin"), newsHandler.GetAllNews)
}

