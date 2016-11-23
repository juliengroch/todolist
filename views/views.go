package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/failures"
	"github.com/juliengroch/todolist/managers"
	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/resources"
)

// TaskListView get all tasks handler
func TaskListView(c *gin.Context) {
	taskList, err := managers.GetTasks(c)

	if err != nil {
		panic(err)
	}

	tr, err := resources.NewTasks(taskList)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tr,
	})
	return
}

// TaskDetailView get one task handler
func TaskDetailView(c *gin.Context) {
	tm := c.MustGet("task").(*models.Task)

	tr, err := resources.NewTask(tm)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, tr)
	return
}

// TaskCreateView create one task
func TaskCreateView(c *gin.Context) {
	newTask := &payloads.Task{}
	c.Bind(newTask)

	// save task
	tm, err := managers.CreateTask(c, newTask)

	if err != nil {
		failures.HandleError(c, err)
		return
	}

	tr, err := resources.NewTask(tm)

	if err != nil {
		failures.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, tr)
}

// TaskUpdateView update one task
func TaskUpdateView(c *gin.Context) {
	task := c.MustGet("task").(*models.Task)

	payload := &payloads.Task{}
	c.Bind(payload)

	// save task
	tm, err := managers.UpdateTask(c, task, payload)

	if err != nil {
		panic(err)
	}

	tr, err := resources.NewTask(tm)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, tr)
}
