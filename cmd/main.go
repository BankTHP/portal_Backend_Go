package main

import (
	"log"
	"os"
	"pccth/portal-blog/config"
	"pccth/portal-blog/routes"
	"strconv"
	"time"

	"pccth/portal-blog/internal/handler"
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
		BodyLimit: 524288000, // 500MB
		ReadTimeout: 600 * time.Second,  // 10 minutes
		WriteTimeout: 600 * time.Second, // 10 minutes
	})

	app.Use(cors.New())
	routes.RegisterRoutes(app, db, authMiddleware)

	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL: "/swagger/doc.json",
		DeepLinking: false,
	}))

	port := viper.GetInt("app.port")
	address := ":" + strconv.Itoa(port)

	// สร้าง upload directory ถ้ายังไม่มี
	uploadDir := viper.GetString("app.upload.upload_dir")
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		log.Fatalf("ไม่สามารถสร้าง upload directory ได้: %v", err)
	}

	// เพิ่ม static route สำหรับไฟล์วิดีโอ
	app.Static("/videos", uploadDir)

	// สร้าง video handler
	videoHandler := handler.NewVideoHandler(uploadDir)

	// เพิ่ม route สำหรับอัปโหลดวิดีโอ
	app.Post("/upload/video", videoHandler.UploadVideo)

	if err := app.Listen(address); err != nil {
		log.Fatalf("ไม่สามารถเริ่มเซิร์ฟเวอร์ได้: %v", err)
	}
}
