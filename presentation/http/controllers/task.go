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

	GetTasks(ctx *fiber.Ctx) error
}

type taskController struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskController(taskUsecase usecase.TaskUsecase) TaskController {
	return &taskController{taskUsecase: taskUsecase}
}

// @Summary Create a new Task
// @Tags Tasks
// @Accept   json
// @Produce  json
// @Param    body        body models.CreateTaskRequest true "Task details"
// @Success 200 {object} models.TaskResponse
// @Failure 400 string error
// @Router /tasks [post]
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

// @Summary Update a Task
// @Tags Tasks
// @Accept   json
// @Produce  json
// @Param    body        body models.UpdateTaskRequest true "Task details"
// @Success 200 {object} models.TaskResponse
// @Failure 400 string error
// @Router /tasks [put]
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

// @Summary Delete a Task
// @Tags Tasks
// @Accept   json
// @Produce  json
// @Param    body        body models.DeleteTaskRequest true "Task ID"
// @Success 200 string OK
// @Failure 400 string error
// @Router /tasks [delete]
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

// @Summary List all Tasks
// @Tags Tasks
// @Accept   json
// @Produce  json
// @Param    status query string false "Task status"
// @Success 200 {object} models.ListTasksResponse
// @Failure 400 {string} error
// @Router /tasks [get]
func (t *taskController) GetTasks(ctx *fiber.Ctx) error {
	var req usecase.ListTasksRequest
	status := ctx.Query("status")
	req.TaskStatus = &status

	tasks, err := t.taskUsecase.GetTasks(ctx.Context(), req)
	if err != nil {
		return err
	}

	response := adaptors.ToListTasksResponse(tasks)
	return ctx.Status(fiber.StatusOK).JSON(response)
}
