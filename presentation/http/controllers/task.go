package controllers

import (
	"task_manager/presentation/adaptors"
	"task_manager/presentation/models"
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
	req := models.CreateTaskRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	taskReq := usecase.CreateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      req.Status,
	}

	task, err := t.taskUsecase.CreateTask(ctx.Context(), taskReq)
	if err != nil {
		return err
	}

	response := adaptors.ToTaskResponse(task)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (t *taskController) UpdateTask(ctx *fiber.Ctx) error {
	req := models.UpdateTaskRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	taskReq := usecase.UpdateTaskRequest{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      req.Status,
	}

	task, err := t.taskUsecase.UpdateTask(ctx.Context(), taskReq)
	if err != nil {
		return err
	}

	response := adaptors.ToTaskResponse(task)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (t *taskController) DeleteTask(ctx *fiber.Ctx) error {
	req := models.DeleteTaskRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return FiberFailedBodyParseError(err)
	}

	taskReq := usecase.DeleteTaskRequest{
		ID: req.ID,
	}

	err := t.taskUsecase.DeleteTask(ctx.Context(), taskReq)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString("OK")
}

func (t *taskController) ListTasks(ctx *fiber.Ctx) error {
	var req usecase.ListTasksRequest
	status := ctx.Query("status")
	req.TaskStatus = &status

	tasks, err := t.taskUsecase.ListTasks(ctx.Context(), req)
	if err != nil {
		return err
	}

	response := adaptors.ToListTasksResponse(tasks)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
