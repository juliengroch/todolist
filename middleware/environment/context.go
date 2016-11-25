package environment

import (
	"context"

	"github.com/juliengroch/todolist/constants"
	"github.com/juliengroch/todolist/models"
)

// AuthenticatedUser returns authenticated User instance from Gin context.
func AuthenticatedUser(ctx context.Context) *models.User {
	if v, ok := ctx.Value(constants.AuthUserKey).(*models.User); ok {
		return v
	}
	return nil
}
