package managers

import (
	"context"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/sanitizing"
	"github.com/juliengroch/todolist/store"
)

// GetTaskByID get one task by id
func GetTaskByID(ctx context.Context, id string, userID string) (*models.Task, error) {
	return store.GetTaskByID(ctx, id, userID)
}

// FindTasks get all tasks
func FindTasks(ctx context.Context, userID string) ([]models.Task, error) {
	return store.FindTasks(ctx, userID)
}

// CreateTask create a task
func CreateTask(ctx context.Context, payload *payloads.Task) (*models.Task, error) {
	sanitizer := sanitizing.FromContext(ctx)

	task := &models.Task{
		ID:          strings.Replace(uuid.NewV4().String(), "-", "", -1),
		Title:       sanitizer.Sanitize(payload.Title),
		Description: sanitizer.Sanitize(payload.Description),
		Priority:    payload.Priority,
		UserID:      payload.User.ID,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	err := store.Create(ctx, task)

	task.User = *payload.User

	return task, err
}

// UpdateTask update a task
func UpdateTask(ctx context.Context, task *models.Task, payload *payloads.Task) (*models.Task, error) {

	task.Title = payload.Title
	task.Description = payload.Description
	task.Priority = payload.Priority
	task.Modified = time.Now()

	return task, store.Save(ctx, task)
}
