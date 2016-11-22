package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"taskProject/constants"
	"taskProject/managers"
	"taskProject/payloads"
)

func main() {
	var server = gin.Default()
	server.GET("/tasks", func(c *gin.Context) {
		taskList, ok := c.Get(constants.TaskListKey)

		if ok {
			c.JSON(http.StatusOK, taskList)
			return
		}

		c.Status(http.StatusNoContent)
	})

	server.POST("/tasks", func(c *gin.Context) {
		newTask := payloads.Task{ID: 1}
		c.Bind(&newTask)

		// save task
		modelTask, err := managers.CreateTasks(c, &newTask)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, modeltask)
	})

	server.Run() // listen and server on 0.0.0.0:8080
}
