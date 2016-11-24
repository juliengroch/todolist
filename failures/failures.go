package failures

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mholt/binding"
)

// HandleError is the main handler for errors
func HandleError(c *gin.Context, err error) {

	switch errv := err.(type) {

	// all errors from unvalid body request field
	case binding.Errors:
		httpError := ValidationError(errv)
		c.JSON(httpError.HTTPStatus, httpError)

	// all GORM errors
	case gorm.Errors:
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error",
	})
}
