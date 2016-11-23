package server

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/middleware"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/views"
)

func Run(ctx context.Context) error {
	var server = gin.Default()

	server.Use(middleware.SetStore(store.FromContext(ctx)))
	server.Use(middleware.SetConfig(config.FromContext(ctx)))

	server.GET("/tasks/:id", views.TaskView)
	server.POST("/tasks", views.TaskCreateView)
	server.PATCH("/tasks/:id", views.TaskUpdateView)

	server.Run() // listen and server on 0.0.0.0:8080

	return nil
}
