package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mholt/binding"

	"github.com/juliengroch/todolist/constants"
	"github.com/juliengroch/todolist/failures"
	"github.com/juliengroch/todolist/managers"
	"github.com/juliengroch/todolist/middleware/environment"
	"github.com/juliengroch/todolist/models"
	"github.com/juliengroch/todolist/payloads"
	"github.com/juliengroch/todolist/resources"
)

// TaskListView get all tasks handler
func TaskListView(c *gin.Context) {
	taskList, err := managers.FindTasks(c, environment.AuthenticatedUser(c).ID)
	if err != nil {
		failures.HandleError(c, err)
	}

	tr, err := resources.NewTasks(taskList)
	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tr,
	})
	return
}

// TaskDetailView get one task handler
func TaskDetailView(c *gin.Context) {
	tm := c.MustGet(constants.TaskKey).(*models.Task)

	tr, err := resources.NewTask(tm)
	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, tr)
	return
}

// TaskCreateView create one task
func TaskCreateView(c *gin.Context) {
	newTask := &payloads.Task{}

	errs := binding.Bind(c.Request, newTask)
	if errs.Len() > 0 {
		failures.HandleError(c, errs)
	}

	// save task
	newTask.User = environment.AuthenticatedUser(c)
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
	task := c.MustGet(constants.TaskKey).(*models.Task)

	payload := &payloads.Task{}

	errs := binding.Bind(c.Request, payload)
	if errs.Len() > 0 {
		failures.HandleError(c, errs)
	}

	// save task
	tm, err := managers.UpdateTask(c, task, payload)
	if err != nil {
		failures.HandleError(c, err)
	}

	tr, err := resources.NewTask(tm)
	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, tr)
}
