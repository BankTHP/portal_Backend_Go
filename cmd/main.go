package main

import (
	"log"
	"pccth/portal-blog/config"
	"pccth/portal-blog/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	_ "pccth/portal-blog/cmd/docs"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//	@title			API สำหรับบล็อกพอร์ทัล
//	@version		1.0
//	@description	นี่คือ API สำหรับบล็อกพอร์ทัล
//	@host			localhost:8080
//	@BasePath		/
func main() {
	config.InitConfig()

	db := config.InitDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(cors.New())
	routes.RegisterRoutes(app, db)

	port := viper.GetInt("app.port")
	address := ":" + strconv.Itoa(port)

	log.Printf("เซิร์ฟเวอร์กำลังทำงานที่พอร์ต %d\n", port)

	if err := app.Listen(address); err != nil {
		log.Fatalf("ไม่สามารถเริ่มเซิร์ฟเวอร์ได้: %v", err)
	}
}
