package resources

import (
	"github.com/juliengroch/todolist/models"
	"github.com/ulule/deepcopier"
)

// Task payload for POST (create)
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
}

// NewTask return Task resource from a Task model instance.
func NewTask(task *models.Task) (*Task, error) {
	resource := &Task{}

	if err := deepcopier.Copy(task).To(resource); err != nil {
		return nil, err
	}

	return resource, nil
}
