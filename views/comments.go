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

// CommentView get one Comment handler
func CommentView(c *gin.Context) {
	tm := c.MustGet(constants.CommentKey).(*models.Comment)

	tr, err := resources.NewComment(tm)

	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, tr)
	return
}

// CommentView get one Comment handler
func UserCommentListView(c *gin.Context) {
	commentList, err := managers.FindComments(c, environment.AuthenticatedUser(c).ID)

	if err != nil {
		failures.HandleError(c, err)
	}

	clr, err := resources.NewComments(commentList)

	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": clr,
	})
	return
}

// CommentCreateView create one Comment
func CommentCreateView(c *gin.Context) {
	newComment := &payloads.Comment{}

	errs := binding.Bind(c.Request, newComment)
	if errs.Len() > 0 {
		failures.HandleError(c, errs)
	}

	newComment.TaskID = c.Param("id")

	// save Comment
	newComment.User = environment.AuthenticatedUser(c)
	tm, err := managers.CreateComment(c, newComment)

	if err != nil {
		failures.HandleError(c, err)
		return
	}

	tr, err := resources.NewComment(tm)

	if err != nil {
		failures.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, tr)
}

// CommentUpdateView update one Comment
func CommentUpdateView(c *gin.Context) {
	comment := c.MustGet(constants.CommentKey).(*models.Comment)

	payload := &payloads.Comment{}

	errs := binding.Bind(c.Request, payload)
	if errs.Len() > 0 {
		failures.HandleError(c, errs)
	}

	// save comment
	tm, err := managers.UpdateComment(c, comment, payload)

	if err != nil {
		failures.HandleError(c, err)
	}

	tr, err := resources.NewComment(tm)

	if err != nil {
		failures.HandleError(c, err)
	}

	c.JSON(http.StatusOK, tr)
}
