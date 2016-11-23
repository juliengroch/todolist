package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/managers"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/resources"
)

// TaskListView get all tasks handler
func TaskListView(c *gin.Context) {
	taskList, err := managers.GetTasks(c)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, taskList)
	return
}

// TaskView get one task handler
func TaskView(c *gin.Context) {
	tm, err := managers.GetOneTask(c, c.Param("id"))

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
	tm, err := managers.CreateTasks(c, newTask)

	tr, err := resources.NewTask(tm)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, tr)
}

// TaskUpdateView update one task
func TaskUpdateView(c *gin.Context) {
	upTask := &payloads.Task{}
	c.Bind(upTask)

	// save task
	tm, err := managers.UpdateTasks(c, c.Param("id"), upTask)

	tr, err := resources.NewTask(tm)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, tr)
}
