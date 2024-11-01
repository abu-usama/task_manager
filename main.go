package main

import (
	"fmt"
	"log"

	"task_manager/config"
	"task_manager/database"
	"task_manager/docs"
	"task_manager/presentation/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title TASK MANAGER API
// @version 1.0
// @BasePath /api/v1/
func main() {

	log.Println("Go Lang Server Running")

	conf := config.GetEnv()
	docs.SwaggerInfo.Host = conf.SwaggerHost
	corsConfig := config.CorsConfig()

	fmt.Println(conf)

	database.ConnectDb()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(cors.New(corsConfig))
	app.Use(recover.New())

	http.Router(app)

	port := conf.Port
	if port == "" {
		port = "3000"
	}

	// Start the server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
