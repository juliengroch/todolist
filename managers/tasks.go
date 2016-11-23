package managers

import (
	"context"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/store"
)

//GetOneTask get one task by id
func GetOneTask(ctx context.Context, id string) (*models.Task, error) {
	return store.FromContext(ctx).GetTaskByID(id)
}

//GetTasks get all task
func GetTasks(ctx context.Context) ([]models.Task, error) {
	return nil, nil
}

// CreateTasks create a task
func CreateTasks(ctx context.Context, payload *payloads.Task) (*models.Task, error) {
	return store.FromContext(ctx).CreateTask(payload.Title, payload.Description, payload.Priority)
}

// UpdateTasks update a task
func UpdateTasks(ctx context.Context, id string, payload *payloads.Task) (*models.Task, error) {
	return store.FromContext(ctx).UpdateTask(id, payload.Title, payload.Description, payload.Priority)
}
