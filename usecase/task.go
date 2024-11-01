package usecase

import (
	"context"
	"errors"
	"time"

	"task_manager/domain/entity"
	"task_manager/domain/task"
	taskoptions "task_manager/domain/task/task_option"

	"github.com/go-playground/validator/v10"
)

type CreateTaskRequest struct {
	Title       string     `validate:"required"`
	Description string     `validate:"required"`
	DueDate     *time.Time `validate:"omitempty"`
	Status      string     `validate:"required,uppercase"`
}

func (r CreateTaskRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

type UpdateTaskRequest struct {
	ID          int        `validate:"required"`
	Title       string     `validate:"required"`
	Description string     `validate:"required"`
	DueDate     *time.Time `validate:"omitempty"`
	Status      string     `validate:"required,uppercase"`
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

	GetTasks(ctx context.Context, req ListTasksRequest) (entity.Tasks, error)
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

	status, exists := entity.StringToTaskStatusMapping[req.Status]
	if !exists {
		return nil, errors.New("invalid status value")
	}

	task := &entity.Task{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      status,
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

	status, exists := entity.StringToTaskStatusMapping[req.Status]
	if !exists {
		return nil, errors.New("invalid status value")
	}

	task := &entity.Task{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		DueDate:     req.DueDate,
		Status:      status,
		UpdatedAt:   time.Now(),
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

func (t *taskUsecase) GetTasks(ctx context.Context, req ListTasksRequest) (entity.Tasks, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var options []taskoptions.OptionSetter
	if req.TaskStatus != nil {
		if status, exists := entity.StringToTaskStatusMapping[*req.TaskStatus]; exists {
			options = append(options, taskoptions.WithStatus(status))
		}
	}

	var tasks entity.Tasks
	var err error
	if len(options) > 0 {
		tasks, err = t.taskRepo.Get(ctx, options...)
	} else {
		tasks, err = t.taskRepo.Get(ctx)
	}
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
