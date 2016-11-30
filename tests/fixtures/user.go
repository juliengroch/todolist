package fixtures

import (
	"context"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// users test info
const (
	UserTestID = "123"
	Username   = "user_test"
	APIKey     = "user_key"

	OtherUserID     = "otheruser123456"
	OtherUserName   = "other_user"
	OtherUserAPIKey = "other_user_key"
)

// AddUserTest add a test user to bdd
func AddUserTest(ctx context.Context) error {
	user := &models.User{
		ID:       UserTestID,
		Username: Username,
		APIKey:   APIKey,
	}

	user2 := &models.User{
		ID:       OtherUserID,
		Username: OtherUserName,
		APIKey:   OtherUserAPIKey,
	}

	err := store.Create(ctx, user)
	if err != nil {
		return err
	}

	return store.Create(ctx, user2)
}
