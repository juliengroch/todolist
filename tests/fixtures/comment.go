package fixtures

import (
	"context"
	"time"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// CommentTestID id of test comment
const CommentTestID = "789"

// AddCommentTest add a test comment to bdd
func AddCommentTest(ctx context.Context) error {
	comment := &models.Comment{
		ID:       CommentTestID,
		Message:  "a test comment",
		UserID:   UserTestID,
		TaskID:   TaskTestID,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	return store.FromContext(ctx).Create(comment)
}
