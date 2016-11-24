package managers

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/payloads/validators"
	"github.com/juliengroch/todolist/store"
)

//GetTaskByID get one task by id
func GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	return store.FromContext(ctx).GetTaskByID(id)
}

//GetTasks get all task
func FindTasks(ctx context.Context) ([]models.Task, error) {
	return store.FromContext(ctx).FindTasks()
}

// CreateTask create a task
func CreateTask(ctx context.Context, payload *payloads.Task) (*models.Task, error) {
	result, err := validators.Task(payload)

	spew.Dump(err, result)

	if err != nil {
		return nil, err
	}

	return store.FromContext(ctx).CreateTask(payload.Title, payload.Description, payload.Priority)
}

// UpdateTask update a task
func UpdateTask(ctx context.Context, task *models.Task, payload *payloads.Task) (*models.Task, error) {

	task.Title = payload.Title
	task.Description = payload.Description
	task.Priority = payload.Priority

	return task, store.FromContext(ctx).Save(task)
}
