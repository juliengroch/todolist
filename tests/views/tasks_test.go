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

// TestTaskUpdateView nominal test
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

// TestTaskUpdateView bad credentials test
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

// TestTaskUpdateView nominal test
func TestTaskCreateView(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {

		title := "faa task"
		desc := "une tache final"
		var prio int8 = 1

		body := map[string]interface{}{
			"title":       title,
			"description": desc,
			"priority":    prio,
		}

		resp := tests.POST(router, &tests.Request{
			URI: "/tasks",
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
			Payload: body,
		})

		assert.Equal(t, http.StatusCreated, resp.Code)

		// test resource
		task := &resources.Task{}
		err := json.Unmarshal(resp.Body.Bytes(), task)
		assert.Nil(t, err)

		assert.NotNil(t, task)
		assert.NotZero(t, task.ID)
		assert.Equal(t, title, task.Title)
		assert.Equal(t, desc, task.Description)
		assert.Equal(t, prio, task.Priority)

		assert.NotNil(t, task.User)
		assert.NotZero(t, task.User.Username)

		// test data in bdd
		tm, err := store.FromContext(ctx).GetTaskByID(task.ID, fixtures.UserTestID)
		assert.Nil(t, err)
		assert.Equal(t, title, tm.Title)
		assert.Equal(t, desc, tm.Description)
		assert.Equal(t, prio, tm.Priority)
	})
}

// TestTaskUpdateView nominal test
func TestTaskUpdateView(t *testing.T) {
	tests.Runner(t, func(router *gin.Engine, ctx context.Context) {

		title := "foo task"
		desc := "update description"
		var prio int8 = 2

		body := map[string]interface{}{
			"title":       title,
			"description": desc,
			"priority":    prio,
		}

		resp := tests.PATCH(router, &tests.Request{
			URI: fmt.Sprintf("/tasks/%s", fixtures.TaskTestID),
			Auth: &tests.Auth{
				Name: fixtures.Username,
				Key:  fixtures.APIKey,
			},
			Payload: body,
		})

		assert.Equal(t, http.StatusOK, resp.Code)

		// test resource
		task := &resources.Task{}
		err := json.Unmarshal(resp.Body.Bytes(), task)
		assert.Nil(t, err)

		assert.NotNil(t, task)
		assert.Equal(t, fixtures.TaskTestID, task.ID)
		assert.Equal(t, title, task.Title)
		assert.Equal(t, desc, task.Description)
		assert.Equal(t, prio, task.Priority)

		assert.NotNil(t, task.User)
		assert.NotZero(t, task.User.Username)

		// test data in bdd
		tm, err := store.FromContext(ctx).GetTaskByID(task.ID, fixtures.UserTestID)
		assert.Nil(t, err)
		assert.Equal(t, title, tm.Title)
		assert.Equal(t, desc, tm.Description)
		assert.Equal(t, prio, tm.Priority)
	})
}
