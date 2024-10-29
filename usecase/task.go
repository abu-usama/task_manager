package usecase

import (
	"context"
	"time"

	"task_manager/domain/entity"
	"task_manager/domain/task"
	taskoptions "task_manager/domain/task/task_option"

	"github.com/go-playground/validator/v10"
)

type CreateTaskRequest struct {
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	DueDate     time.Time `validate:"required"`
	Status      string    `validate:"required,lowercase"`
}

func (r CreateTaskRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type UpdateTaskRequest struct {
	Title       string    `validate:"required"`
	Description string    `validate:"required"`
	DueDate     time.Time `validate:"required"`
	Status      string    `validate:"required,lowercase"`
}

func (r UpdateTaskRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type DeleteTaskRequest struct {
	ID int `validate:"required"`
}

func (r DeleteTaskRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type ListTasksRequest struct {
	TaskStatus *string `validate:"omitempty"`
}

func (r ListTasksRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type TaskUsecase interface {
	CreateTask(ctx context.Context, req CreateTaskRequest) (*entity.Task, error)
	UpdateTask(ctx context.Context, req UpdateTaskRequest) (*entity.Task, error)
	DeleteTask(ctx context.Context, req DeleteTaskRequest) error

	ListTasks(ctx context.Context, req ListTasksRequest) ([]*entity.Task, error)
}

type taskUsecase struct {
	taskRepo task.Repository
}

func NewTaskUsecase(taskRepo task.Repository) TaskUsecase {
	return &taskUsecase{taskRepo}
}

func (t *taskUsecase) CreateTask(ctx context.Context, req CreateTaskRequest) (*entity.Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      entity.StringToStatusTypeMapping[req.Status],
	}

	task, err := t.taskRepo.Create(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskUsecase) UpdateTask(ctx context.Context, req UpdateTaskRequest) (*entity.Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      entity.StringToStatusTypeMapping[req.Status],
	}

	task, err := t.taskRepo.Update(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskUsecase) DeleteTask(ctx context.Context, req DeleteTaskRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	err := t.taskRepo.Delete(ctx, req.ID)
	if err != nil {
		return err
	}

	return nil
}

func (t *taskUsecase) ListTasks(ctx context.Context, req ListTasksRequest) ([]*entity.Task, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var options taskoptions.OptionSetter
	if req.TaskStatus != nil {
		options = taskoptions.WithStatus(entity.StringToStatusTypeMapping[*req.TaskStatus])
	}

	tasks, err := t.taskRepo.List(ctx, options)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
