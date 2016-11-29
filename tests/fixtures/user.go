package fixtures

import (
	"context"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// user test info
const (
	UserTestID = "123"
	Username   = "user_test"
	APIKey     = "user_key"
)

// AddUserTest add a test user to bdd
func AddUserTest(ctx context.Context) error {
	user := &models.User{
		ID:       UserTestID,
		Username: Username,
		APIKey:   APIKey,
	}

	return store.FromContext(ctx).Create(user)
}
