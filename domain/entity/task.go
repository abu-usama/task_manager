package entity

import "time"

type TaskStatus int

const (
	TaskStatusComplete TaskStatus = iota
	TaskStatusPending
	TaskStatusOngoing
)

var StringToTaskStatusMapping = map[string]TaskStatus{
	"COMPLETE": TaskStatusComplete,
	"PENDING":  TaskStatusPending,
	"ONGOING":  TaskStatusOngoing,
}

var TaskStatusToStringMapping = map[TaskStatus]string{
	TaskStatusComplete: "COMPLETE",
	TaskStatusPending:  "PENDING",
	TaskStatusOngoing:  "ONGOING",
}

type Tasks []Task

type Task struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	Status      TaskStatus
	DueDate     *time.Time
}
