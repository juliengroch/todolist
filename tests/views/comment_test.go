package views

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/juliengroch/todolist/middleware"
	"github.com/juliengroch/todolist/resources"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/tests"
	"github.com/juliengroch/todolist/tests/fixtures"
)

// TestCommentUpdateView nominal test
func TestCommentDetailViewAuthOK(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {
		resp := tests.GET(router, &tests.Request{
			URI: fmt.Sprintf("/comments/%s", fixtures.CommentTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
		})

		assert.Equal(t, http.StatusOK, resp.Code)

		comment := &resources.Comment{}
		err := json.Unmarshal(resp.Body.Bytes(), comment)
		assert.Nil(t, err)

		assert.NotNil(t, comment)
		assert.NotZero(t, comment.ID)
		assert.NotZero(t, comment.Message)
		assert.NotNil(t, comment.User)
		assert.NotZero(t, comment.User.Username)
	})
}

// TestCommentUpdateView bad credentials test
func TestCommentDetailViewAuthNOK(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {
		resp := tests.GET(router, &tests.Request{
			URI: fmt.Sprintf("/comments/%s", fixtures.CommentTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  "bad_key",
			},
		})

		respError := &middleware.Error{}
		err := json.Unmarshal(resp.Body.Bytes(), respError)
		assert.Nil(t, err)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.NotNil(t, respError)
		assert.Equal(t, middleware.ErrInvalidCredentials.Message, respError.Message)
	})
}

// TestCommentUpdateView nominal test
func TestCommentCreateView(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {

		mess := "a test comment"

		body := map[string]interface{}{
			"message": mess,
		}

		resp := tests.POST(router, &tests.Request{
			URI: fmt.Sprintf("/tasks/%s/comments", fixtures.TaskTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
			Payload: body,
		})

		assert.Equal(t, http.StatusCreated, resp.Code)

		// test resource
		comment := &resources.Comment{}
		err := json.Unmarshal(resp.Body.Bytes(), comment)
		assert.Nil(t, err)

		assert.NotNil(t, comment)
		assert.NotZero(t, comment.ID)
		assert.Equal(t, mess, comment.Message)
		assert.Equal(t, fixtures.TaskTestID, comment.TaskID)
		assert.NotNil(t, comment.User)
		assert.Equal(t, fixtures.Username, comment.User.Username)

		// test data in bdd
		tm, err := store.GetCommentByID(ctx, comment.ID, fixtures.UserTestID)
		assert.Nil(t, err)
		assert.Equal(t, mess, tm.Message)
		assert.Equal(t, fixtures.TaskTestID, tm.TaskID)
	})
}

// TestCommentUpdateView nominal test
func TestCommentUpdateView(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {

		mess := "update comment message"

		body := map[string]interface{}{
			"message": mess,
		}

		resp := tests.PATCH(router, &tests.Request{
			URI: fmt.Sprintf("/comments/%s", fixtures.CommentTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
			Payload: body,
		})

		assert.Equal(t, http.StatusOK, resp.Code)

		// test resource
		comment := &resources.Comment{}
		err := json.Unmarshal(resp.Body.Bytes(), comment)
		assert.Nil(t, err)

		assert.NotNil(t, comment)
		assert.Equal(t, fixtures.CommentTestID, comment.ID)
		assert.Equal(t, mess, comment.Message)
		assert.Equal(t, fixtures.TaskTestID, comment.TaskID)
		assert.NotNil(t, comment.User)
		assert.Equal(t, fixtures.Username, comment.User.Username)

		// test data in bdd
		tm, err := store.GetCommentByID(ctx, comment.ID, fixtures.UserTestID)
		assert.Nil(t, err)
		assert.Equal(t, mess, tm.Message)
		assert.Equal(t, fixtures.TaskTestID, tm.TaskID)
	})
}

// TestUpNoOwnerComment test if task owner can edit task comment from other user
func TestUpNoOwnerComment(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {

		mess := "update comment message from other user on his task"

		body := map[string]interface{}{
			"message": mess,
		}

		resp := tests.PATCH(router, &tests.Request{
			URI: fmt.Sprintf("/comments/%s", fixtures.CommentNoOwnerID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
			Payload: body,
		})

		assert.Equal(t, http.StatusOK, resp.Code)

		// test resource
		comment := &resources.Comment{}
		err := json.Unmarshal(resp.Body.Bytes(), comment)
		assert.Nil(t, err)

		assert.NotNil(t, comment)
		assert.Equal(t, fixtures.CommentNoOwnerID, comment.ID)
		assert.Equal(t, mess, comment.Message)
		assert.Equal(t, fixtures.TaskTestID, comment.TaskID)
		assert.NotNil(t, comment.User)
		assert.Equal(t, fixtures.OtherUserName, comment.User.Username)

		// test data in bdd
		tm, err := store.GetCommentByID(ctx, comment.ID, fixtures.UserTestID)
		assert.Nil(t, err)
		assert.Equal(t, mess, tm.Message)
		assert.Equal(t, fixtures.TaskTestID, tm.TaskID)
	})
}
