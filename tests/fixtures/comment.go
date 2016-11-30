package fixtures

import (
	"context"
	"time"

	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/store"
)

// CommentTestID id of test comment
const CommentTestID = "789"

// CommentNoOwnerID id of test comment from no owner task user
const CommentNoOwnerID = "123648289"

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

	comment2 := &models.Comment{
		ID:       CommentNoOwnerID,
		Message:  "comment from other user",
		UserID:   OtherUserID,
		TaskID:   TaskTestID,
		Created:  time.Now(),
		Modified: time.Now(),
	}

	err := store.Create(ctx, comment)
	if err != nil {
		return err
	}

	return store.Create(ctx, comment2)
}
