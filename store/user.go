package store

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/juliengroch/todolist/models"
)

// GetUserByUsernameAndAPIKey returns a User instance for a given username and API Key.
func GetUserByUsernameAndAPIKey(ctx context.Context, username string, key string) (*models.User, error) {
	user, err := FromContext(ctx).GetUserByKey(key)
	if err != nil {
		return nil, err
	}

	if user.Username != username {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}
