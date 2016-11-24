package failures

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// HandleError is the main handler for errors
func HandleError(c *gin.Context, err error) {
	spew.Dump(err)
	switch err.(type) {

	// all errors from unvalid body request field
	case govalidator.Errors:
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return

	case gorm.Errors:
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error",
	})
}
