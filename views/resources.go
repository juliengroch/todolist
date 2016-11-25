package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/juliengroch/todolist/constants"
	"github.com/juliengroch/todolist/managers"
)

func TaskResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		tm, err := managers.GetTaskByID(c, c.Param("id"))

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.AbortWithStatus(http.StatusNotFound)
			}
		}

		c.Set(constants.TaskKey, tm)

		c.Next()
	}
}
