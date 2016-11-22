package managers

import (
	"context"

	"taskProject/models"
	"taskProject/payloads"
)

// CreateTasks create a task
func CreateTasks(ctx context.Context, payload *payloads.Task) (models.Task, error) {
	task := models.Task{}
	return task
}
