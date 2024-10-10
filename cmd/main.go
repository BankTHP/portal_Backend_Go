package main

import (
	"log"
	"pccth/portal-blog/config"
	"pccth/portal-blog/routes"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()

	db := config.InitDB()

	app := fiber.New()

	routes.RegisterRoutes(app, db)

    port := viper.GetInt("app.port") 

    if err := app.Listen(":" + strconv.Itoa(port)); err != nil { 
        log.Fatalf("Failed to start the server: %v", err)
    }

}
