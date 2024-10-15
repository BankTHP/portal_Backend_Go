package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "pccth/portal-blog/cmd/docs"
)

func SwaggerRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
	}))
}
