package routes

import (
	"pccth/portal-blog/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func ReleaseNoteRoutes(app *fiber.App, releaseNoteHandler *handler.ReleaseNoteHandler) {
    app.Post("/createReleaseNote", releaseNoteHandler.CreateReleaseNote)             

}

