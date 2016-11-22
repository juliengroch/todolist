package main

import (
	"net/http"
	"taskProject/constants"
	"taskProject/managers"
	"taskProject/payloads"

	"github.com/gin-gonic/gin"
)

func main() {
	var server = gin.Default()
	server.GET("/tasks", func(c *gin.Context) {
		taskList, isOk := c.Get(constants.TaskListKey)

		if isOk {
			c.JSON(http.StatusOK, taskList)

		} else {
			c.Status(http.StatusNoContent)
		}
	})

	server.POST("/tasks", func(c *gin.Context) {
		newTask := payloads.Task{ID: 1}
		c.Bind(&newTask)

		// save task
		modeltask, err := managers.CreateTasks(c, &newTask)
		c.JSON(http.StatusCreated, modeltask)
	})

	server.Run() // listen and server on 0.0.0.0:8080
}
