package main

import (
	"log"
	"pccth/portal-blog/config"
	"pccth/portal-blog/routes"
	"strconv"

	_ "pccth/portal-blog/internal/handler"
	"pccth/portal-blog/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

// @title Portal Blog API
// @version 1.0
// @description API สำหรับระบบบล็อกของพอร์ทัล
// @termsOfService http://swagger.io/terms/
// @contact.name ทีมสนับสนุน API
// @contact.email support@portalblog.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	config.InitConfig()
	if err := config.InitPublicKey(); err != nil {
		log.Fatalf("ไม่สามารถเริ่มต้นคีย์สาธารณะได้: %v", err)
	}
	db := config.InitDB()
	authMiddleware := middleware.NewAuthMiddleware(config.GetPublicKey())
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	routes.RegisterRoutes(app, db, authMiddleware)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
		DeepLinking: false,
	}))

	port := viper.GetInt("app.port")
	address := ":" + strconv.Itoa(port)


	if err := app.Listen(address); err != nil {
		log.Fatalf("ไม่สามารถเริ่มเซิร์ฟเวอร์ได้: %v", err)
	}
}
