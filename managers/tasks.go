package managers

import (
	"context"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/store"
)

// GetTaskByID get one task by id
func GetTaskByID(ctx context.Context, id string, userID string) (*models.Task, error) {
	return store.FromContext(ctx).GetTaskByID(id, userID)
}

// FindTasks get all task
func FindTasks(ctx context.Context) ([]models.Task, error) {
	return store.FromContext(ctx).FindTasks()
}

// CreateTask create a task
func CreateTask(ctx context.Context, payload *payloads.Task) (*models.Task, error) {
	task := &models.Task{
		ID:          strings.Replace(uuid.NewV4().String(), "-", "", -1),
		Title:       payload.Title,
		Description: payload.Description,
		Priority:    payload.Priority,
		UserID:      payload.User.ID,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	err := store.FromContext(ctx).Create(task)

	return task, err
}

// UpdateTask update a task
func UpdateTask(ctx context.Context, task *models.Task, payload *payloads.Task) (*models.Task, error) {

	task.Title = payload.Title
	task.Description = payload.Description
	task.Priority = payload.Priority

	return task, store.FromContext(ctx).Save(task)
}
