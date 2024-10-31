package http

import (
	"task_manager/config"
	"task_manager/database"
	postgress "task_manager/infra/postgres"
	controllers "task_manager/presentation/http/controllers"
	"task_manager/usecase"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	config.GetEnv()

	taskCtrl := controllers.NewTaskController(usecase.NewTaskUsecase(postgress.NewTaskRepositoryPostgres(database.DB)))
	apiV1 := app.Group("/api/v1")

	apiV1.Post("/tasks", taskCtrl.CreateTask)
	apiV1.Put("/tasks", taskCtrl.UpdateTask)
	apiV1.Delete("/tasks", taskCtrl.DeleteTask)
	apiV1.Get("/tasks", taskCtrl.ListTasks)
}
