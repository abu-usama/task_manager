package adaptors

import (
	"task_manager/domain/entity"
	"task_manager/presentation/models"
)

func ToTaskResponse(task *entity.Task) models.Common {
	taskResponse := models.TaskResponse{
		ID:          task.ID,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Status:      entity.TaskStatusToStringMapping[task.Status],
	}

	return models.Common{
		Success: true,
		Data:    taskResponse,
	}
}

func ToListTasksResponse(tasks []*entity.Task) models.Common {
	tasksResponse := models.ListTasksResponse{}
	for _, task := range tasks {
		taskResponse := &models.TaskResponse{
			ID:          task.ID,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Status:      entity.TaskStatusToStringMapping[task.Status],
		}
		tasksResponse.Tasks = append(tasksResponse.Tasks, taskResponse)
	}

	return models.Common{
		Success: true,
		Data:    tasksResponse,
	}
}
