package resources

import (
	"github.com/juliengroch/todolist/models"
	"github.com/ulule/deepcopier"
)

// Comment resource
type Comment struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	TaskID  string `json:"taskid"`
	User    *User  `json:"user"`
}

// NewComment return Comment resource from a Comment model instance.
func NewComment(comment *models.Comment) (*Comment, error) {
	resource := &Comment{}

	if err := deepcopier.Copy(comment).To(resource); err != nil {
		return nil, err
	}

	var err error
	resource.User, err = NewUser(&comment.User)

	if err != nil {
		return nil, err
	}

	return resource, nil
}

// NewComments multiple taks resources
func NewComments(comments []models.Comment) ([]*Comment, error) {
	resources := []*Comment{}

	for _, tm := range comments {
		resource, err := NewComment(&tm)
		if err != nil {
			return nil, err
		}

		resources = append(resources, resource)
	}

	return resources, nil
}
