package routes

import (
	"pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReleaseNoteRoutes(app *fiber.App, releaseNoteHandler *handler.ReleaseNoteHandler, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/releaseNote")
	api.Post("/createReleaseNote", authMiddleware.HasRole("client_user", "client_admin"), 	releaseNoteHandler.CreateReleaseNote)
	api.Get("/getReleaseNoteById/:id", authMiddleware.HasRole("client_user", "client_admin"), releaseNoteHandler.GetReleaseByID)
	api.Put("/updateReleaseNote/:id", authMiddleware.HasRole("client_user", "client_admin"), releaseNoteHandler.UpdateRelease)
	api.Delete("/deleteReleaseNote/:id", authMiddleware.HasRole("client_user", "client_admin"), releaseNoteHandler.DeleteRelease)
	api.Get("/getAllReleaseNotes", authMiddleware.HasRole("client_user", "client_admin"), releaseNoteHandler.GetAllRelease)

}
