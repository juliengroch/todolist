package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/juliengroch/todolist/constants"
	"github.com/juliengroch/todolist/managers"
	"github.com/juliengroch/todolist/middleware/environment"
)

func TaskResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		tm, err := managers.GetTaskByID(c, c.Param("id"), environment.AuthenticatedUser(c).ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}

		c.Set(constants.TaskKey, tm)

		c.Next()
	}
}

func CommentResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		tm, err := managers.GetCommentByID(c, c.Param("id"), environment.AuthenticatedUser(c).ID)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}

		c.Set(constants.CommentKey, tm)

		c.Next()
	}
}
