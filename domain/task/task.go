package task

import (
	"context"

	"task_manager/domain/entity"
	"task_manager/domain/task/task_option"
)

// Repository combines the Read and Write repositories for task data operations.
type Repository interface {
	ReadRepository
	WriteRepository
}

type WriteRepository interface {
	Create(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Update(ctx context.Context, task *entity.Task) (*entity.Task, error)
	Delete(ctx context.Context, id int) error
}

type ReadRepository interface {
	Get(ctx context.Context, opts ...task_option.OptionSetter) (entity.Tasks, error)
}
