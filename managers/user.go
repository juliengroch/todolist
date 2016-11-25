package managers

import (
	"context"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// GetUserByUsernameAndAPIKey returns a User instance for a given username and API Key.
func GetUserByUsernameAndAPIKey(ctx context.Context, username string, key string) (*models.User, error) {
	return store.GetUserByUsernameAndAPIKey(ctx, username, key)
}
