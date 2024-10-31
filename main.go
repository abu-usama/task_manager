package main

import (
	"fmt"
	"log"

	"task_manager/config"
	"task_manager/database"
	"task_manager/presentation/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @title TASK MANAGER API
// @version 1.0
// @BasePath /api/v1/
func main() {

	log.Println("Go Lang Server Running")

	conf := config.GetEnv()

	fmt.Println(conf)
	// swaggerHost := conf.SwaggerHost
	// if swaggerHost == "" {
	// 	swaggerHost = "localhost"
	// }

	// docs.SwaggerInfo.Host = swaggerHost

	database.ConnectDb()

	// errorLogRepo := postgress.NewErrorLogRepositoryPostgres(database.DB)

	// app := fiber.New(fiber.Config{
	// 	ErrorHandler: middlewares.ErrorHandlingMiddleware(errorLogRepo),
	// })

	corsConfig := cors.Config{
		AllowOrigins: "http://locahost:3000,  http://127.0.0.1:3000, http://localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}

	app := fiber.New()
	app.Use(cors.New(corsConfig))
	app.Use(recover.New())
	// app.Use(middlewares.CustomLogger())

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
