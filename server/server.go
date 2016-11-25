package server

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/middleware"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/views"
)

// Run fonction to start the server
func Run(ctx context.Context) error {
	var server = gin.Default()

	server.Use(middleware.SetStore(store.FromContext(ctx)))
	server.Use(middleware.SetConfig(config.FromContext(ctx)))

	taskResource := views.TaskResource()
	auth := middleware.Authentication()

	server.GET("/tasks", auth, views.TaskListView)
	server.GET("/tasks/:id", auth, taskResource, views.TaskDetailView)
	server.POST("/tasks", auth, views.TaskCreateView)
	server.PATCH("/tasks/:id", auth, taskResource, views.TaskUpdateView)

	server.Run()

	return nil
}
