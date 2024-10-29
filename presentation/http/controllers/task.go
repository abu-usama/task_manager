package controllers

import (
	"task_manager/usecase"

	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	GetPost(C *fiber.Ctx) error
}

type taskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(taskUsecase usecase.TaskUsecase) TaskController {
	return &taskController{taskUsecase: taskUsecase}
}

func (t *taskController) GetPost(C *fiber.Ctx) error {
	return nil
}
