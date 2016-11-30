package managers

import (
	"context"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/sanitizing"
	"github.com/juliengroch/todolist/store"
)

// GetCommentByID get one Comment by id
func GetCommentByID(ctx context.Context, id string, userID string) (*models.Comment, error) {
	return store.GetCommentByID(ctx, id, userID)
}

// FindComments get all tasks
func FindComments(ctx context.Context, userID string) ([]models.Comment, error) {
	return store.FindComments(ctx, userID)
}

// CreateComment create a Comment
func CreateComment(ctx context.Context, payload *payloads.Comment) (*models.Comment, error) {
	sanitizer := sanitizing.FromContext(ctx)

	comment := &models.Comment{
		ID:       strings.Replace(uuid.NewV4().String(), "-", "", -1),
		Message:  sanitizer.Sanitize(payload.Message),
		UserID:   payload.User.ID,
		TaskID:   payload.TaskID,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	err := store.Create(ctx, comment)
	if err != nil {
		return nil, err
	}

	comment.User = *payload.User

	return comment, err
}

// UpdateComment update a Comment
func UpdateComment(ctx context.Context, comment *models.Comment, payload *payloads.Comment) (*models.Comment, error) {

	comment.Message = payload.Message
	comment.Modified = time.Now()

	return comment, store.Save(ctx, comment)
}
