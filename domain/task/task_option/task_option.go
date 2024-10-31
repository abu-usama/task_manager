package task_option

import "task_manager/domain/entity"

type TaskOptions struct {
	Status *entity.TaskStatus // The status of a task.
}

type OptionSetter func(*TaskOptions)

func WithStatus(status entity.TaskStatus) OptionSetter {
	return func(opt *TaskOptions) {
		opt.Status = &status
	}
}
