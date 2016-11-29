package views

import (
	"github.com/gin-gonic/gin"

	"github.com/juliengroch/todolist/middleware"
)

// Routes registers all the API urls and views to the router.
func Routes(router *gin.Engine) {
	// middleware
	taskResource := TaskResource()
	commentResource := CommentResource()
	auth := middleware.Authentication()

	// task api
	router.GET("/tasks", auth, TaskListView)
	router.GET("/tasks/:id", auth, taskResource, TaskDetailView)
	router.POST("/tasks", auth, TaskCreateView)
	router.PATCH("/tasks/:id", auth, taskResource, TaskUpdateView)

	// comment api
	router.GET("/comments/:id", auth, commentResource, CommentView)
	router.GET("/users/:id/comments", auth, UserCommentListView)
	router.POST("/tasks/:id/comments", auth, CommentCreateView)
	router.PATCH("/comments/:id", auth, commentResource, CommentUpdateView)
}
