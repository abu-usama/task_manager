package controllers

import (
	"task_manager/presentation/adaptors"
	"task_manager/usecase"

	"github.com/gofiber/fiber/v2"
)

type TaskController interface {
	CreateTask(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error

	ListTasks(ctx *fiber.Ctx) error
}

type taskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(taskUsecase usecase.TaskUsecase) TaskController {
	return &taskController{taskUsecase: taskUsecase}
}

func (t *taskController) CreateTask(ctx *fiber.Ctx) error {
	var req usecase.CreateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	task, err := t.taskUsecase.CreateTask(ctx.Context(), req)
	if err != nil {
		return err
	}

	response := adaptors.ToTaskResponse(task)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (t *taskController) UpdateTask(ctx *fiber.Ctx) error {
	var req usecase.UpdateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	task, err := t.taskUsecase.UpdateTask(ctx.Context(), req)
	if err != nil {
		return err
	}

	response := adaptors.ToTaskResponse(task)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (t *taskController) DeleteTask(ctx *fiber.Ctx) error {
	var req usecase.DeleteTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	err := t.taskUsecase.DeleteTask(ctx.Context(), req)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString("OK")
}

func (t *taskController) ListTasks(ctx *fiber.Ctx) error {
	var req usecase.ListTasksRequest
	// if err := ctx.BodyParser(&req); err != nil {
	// 	return FiberFailedBodyParseError(err)
	// }
	status := ctx.Query("status")
	req.TaskStatus = &status

	tasks, err := t.taskUsecase.ListTasks(ctx.Context(), req)
	if err != nil {
		return err
	}

	response := adaptors.ToListTasksResponse(tasks)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
