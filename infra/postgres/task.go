package postgress

import (
	"context"
	"fmt"
	"time"

	"task_manager/domain/entity"
	"task_manager/domain/task"
	taskoptions "task_manager/domain/task/task_option"

	"gorm.io/gorm"
)

const (
	ErrFailedToGetTasks   = "failed to get tasks %w"
	ErrFailedToCreateTask = "failed to create task: %w"
	ErrFailedToUpdateTask = "failed to update task: %w"
	ErrFailedToDeleteTask = "failed to delete task: %w"
)

type Common struct {
	ID        int       `gorm:"column:primaryKey;column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type TaskRow struct {
	Common
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	Status      int        `gorm:"column:status"`
	DueDate     *time.Time `gorm:"column:due_date"`
}

func (p *TaskRow) TableName() string {
	return "tasks"
}

func (t *TaskRow) ToDomain() *entity.Task {
	return &entity.Task{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		Title:       t.Title,
		Description: t.Description,
		Status:      entity.TaskStatus(t.Status),
		DueDate:     t.DueDate,
	}
}

func (TaskRow) FromDomain(task entity.Task) *TaskRow {
	return &TaskRow{
		Common: Common{
			ID:        task.ID,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		},
		Title:       task.Title,
		Description: task.Description,
		Status:      int(task.Status),
		DueDate:     task.DueDate,
	}
}

type TaskReadRepositoryPostgres struct {
	DB *gorm.DB
}

type TaskWriteRepositoryPostgres struct {
	DB *gorm.DB
}

type TaskRepositoryPostgres struct {
	TaskReadRepositoryPostgres
	TaskWriteRepositoryPostgres
}

func NewTaskRepositoryPostgres(DB *gorm.DB) task.Repository {
	return &TaskRepositoryPostgres{
		TaskReadRepositoryPostgres:  TaskReadRepositoryPostgres{DB},
		TaskWriteRepositoryPostgres: TaskWriteRepositoryPostgres{DB},
	}
}

func (c *TaskWriteRepositoryPostgres) Create(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	taskRow := TaskRow{}.FromDomain(*task)
	if err := c.DB.WithContext(ctx).Create(&taskRow).Error; err != nil {
		return nil, fmt.Errorf(ErrFailedToCreateTask, err)
	}

	return taskRow.ToDomain(), nil
}

func (c *TaskWriteRepositoryPostgres) Update(ctx context.Context, task *entity.Task) (*entity.Task, error) {
	taskRow := TaskRow{}.FromDomain(*task)
	if err := c.DB.WithContext(ctx).Model(&TaskRow{}).Where("id = ?", taskRow.ID).Updates(&taskRow).Error; err != nil {
		return nil, fmt.Errorf(ErrFailedToUpdateTask, err)
	}

	return taskRow.ToDomain(), nil
}

func (c *TaskWriteRepositoryPostgres) Delete(ctx context.Context, id int) error {
	var taskRow TaskRow
	taskRow.ID = id
	if err := c.DB.WithContext(ctx).Delete(&taskRow).Error; err != nil {
		return fmt.Errorf(ErrFailedToDeleteTask, err)
	}

	return nil
}

func (c *TaskWriteRepositoryPostgres) List(ctx context.Context, opts ...taskoptions.OptionSetter) ([]*entity.Task, error) {
	var options = &taskoptions.TaskOptions{}
	for _, setter := range opts {
		setter(options)
	}
	query := c.DB.WithContext(ctx).Model(&TaskRow{}).Order("created_at DESC")

	var taskRows []TaskRow
	if options.Status != nil {
		if int(*options.Status) < 3 {
			query = query.Where("status = ?", int(*options.Status))
		}
	}

	err := query.Find(&taskRows).Error
	if err != nil {
		return nil, err
	}
	var tasks []*entity.Task
	for _, row := range taskRows {
		task := row.ToDomain()
		tasks = append(tasks, task)
	}

	return tasks, nil
}
