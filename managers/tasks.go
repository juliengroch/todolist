package managers

import (
	"context"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
)

// CreateTasks create a task
func CreateTasks(ctx context.Context, payload *payloads.Task) (models.Task, error) {
	var err error
	task := models.Task{}

	return task, err
}
