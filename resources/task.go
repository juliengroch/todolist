package resources

import (
	"github.com/juliengroch/todolist/models"
	"github.com/ulule/deepcopier"
)

// Task resource
type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int8   `json:"priority"`
	User        *User  `json:"user"`
}

// NewTask return Task resource from a Task model instance.
func NewTask(task *models.Task) (*Task, error) {
	resource := &Task{}

	if err := deepcopier.Copy(task).To(resource); err != nil {
		return nil, err
	}

	var err error
	resource.User, err = NewUser(&task.User)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

// NewTasks multiple taks resources
func NewTasks(tasks []models.Task) ([]*Task, error) {
	resources := []*Task{}

	for _, tm := range tasks {
		resource, err := NewTask(&tm)
		if err != nil {
			return nil, err
		}

		resources = append(resources, resource)
	}

	return resources, nil
}
