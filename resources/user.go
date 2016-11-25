package resources

import (
	"github.com/juliengroch/todolist/models"
	"github.com/ulule/deepcopier"
)

// User resource
type User struct {
	Username string `json:"username"`
}

// NewUser return user resource from a user model instance.
func NewUser(user *models.User) (*User, error) {
	resource := &User{}

	if err := deepcopier.Copy(user).To(resource); err != nil {
		return nil, err
	}

	return resource, nil
}
