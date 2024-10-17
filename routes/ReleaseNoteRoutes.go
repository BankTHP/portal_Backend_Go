package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func ReleaseNoteRoutes(app *fiber.App, releaseNoteHandler *handler.ReleaseNoteHandler) {
	api := app.Group("/releaseNote")
	api.Post("/createReleaseNote", releaseNoteHandler.CreateReleaseNote)
	api.Get("/getReleaseNoteById/:id", releaseNoteHandler.GetReleaseByID)
	api.Put("/updateReleaseNote/:id", releaseNoteHandler.UpdateRelease)
	api.Delete("/deleteReleaseNote/:id", releaseNoteHandler.DeleteRelease)
	api.Get("/getAllReleaseNotes", releaseNoteHandler.GetAllRelease)

}
