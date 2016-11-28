package server

import (
	"context"

	"github.com/gin-gonic/gin"

	"fmt"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/middleware"
	"github.com/juliengroch/todolist/sanitizing"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/views"
)

// Run fonction to start the server
func Run(ctx context.Context) error {
	var server = gin.Default()

	cfg := config.FromContext(ctx)

	appContext := &middleware.ApplicationContextOptions{
		Config:    cfg,
		Store:     store.FromContext(ctx),
		Logger:    loggers.FromContext(ctx),
		Sanitizer: sanitizing.FromContext(ctx),
	}

	server.Use(middleware.ApplicationContext(appContext))

	taskResource := views.TaskResource()
	commentResource := views.CommentResource()
	auth := middleware.Authentication()

	// task api
	server.GET("/tasks", auth, views.TaskListView)
	server.GET("/tasks/:id", auth, taskResource, views.TaskDetailView)
	server.POST("/tasks", auth, views.TaskCreateView)
	server.PATCH("/tasks/:id", auth, taskResource, views.TaskUpdateView)

	// comment api
	server.GET("/comments/:id", auth, commentResource, views.CommentView)
	server.GET("/users/:id/comments", auth, views.UserCommentListView)
	server.POST("/tasks/:id/comments", auth, views.CommentCreateView)
	server.PATCH("/comments/:id", auth, commentResource, views.CommentUpdateView)

	server.Run(fmt.Sprintf(":%d", cfg.Server.Port))

	return nil
}
