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
	"github.com/juliengroch/todolist/tests"
	"github.com/juliengroch/todolist/tests/fixtures"
)

func TestTaskDetailViewAuthOK(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {
		resp := tests.GET(router, &tests.Request{
			URI: fmt.Sprintf("/tasks/%s", fixtures.TaskTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
		})

		assert.Equal(t, http.StatusOK, resp.Code)

		task := &resources.Task{}
		err := json.Unmarshal(resp.Body.Bytes(), task)
		assert.Nil(t, err)

		assert.NotNil(t, task)
		assert.NotZero(t, task.ID)
		assert.NotZero(t, task.Title)
		assert.NotZero(t, task.Description)
		assert.NotNil(t, task.User)
		assert.NotZero(t, task.User.Username)
		assert.NotNil(t, task.Comments)
	})
}

func TestTaskDetailViewAuthNOK(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {
		resp := tests.GET(router, &tests.Request{
			URI: fmt.Sprintf("/tasks/%s", fixtures.TaskTestID),
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
