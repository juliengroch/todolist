package fixtures

import (
	"context"
	"time"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// TaskTestID id of test task
const TaskTestID = "456"

// AddTaskTest add a test task to bdd
func AddTaskTest(ctx context.Context) error {
	task := &models.Task{
		ID:          TaskTestID,
		Title:       "foo task",
		Description: "a test task",
		Priority:    1,
		UserID:      UserTestID,
		Created:     time.Now(),
		Modified:    time.Now(),
	}

	return store.FromContext(ctx).Create(task)
}
