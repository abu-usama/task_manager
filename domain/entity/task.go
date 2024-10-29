package entity

import "time"

type StatusType int

const (
	StatusTypeComplete StatusType = iota
	StatusTypePending
	StatusTypeOngoing
)

var StringToStatusTypeMapping = map[string]StatusType{
	"complete": StatusTypeComplete,
	"pending":  StatusTypePending,
	"ongoing":  StatusTypeOngoing,
}

var StatusTypeToStringMapping = map[StatusType]string{
	StatusTypeComplete: "complete",
	StatusTypePending:  "pending",
	StatusTypeOngoing:  "ongoing",
}

type Task struct {
	ID          int
	Title       string
	Description string
	Status      StatusType
	DueDate     time.Time
}
