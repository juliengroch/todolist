package resources

import (
	"github.com/juliengroch/todolist/models"
	"github.com/ulule/deepcopier"
)

// Task resource
type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    int8       `json:"priority"`
	User        *User      `json:"user"`
	Comments    []*Comment `json:"comments"`
}

// NewTask return Task resource from a Task model instance.
func NewTask(task *models.Task) (*Task, error) {
	resource := &Task{}
	var err error

	err = deepcopier.Copy(task).To(resource)

	if err != nil {
		return nil, err
	}

	// map with user
	resource.User, err = NewUser(&task.User)

	if err != nil {
		return nil, err
	}

	// map with Comments
	resource.Comments, err = NewComments(task.Comments)

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
