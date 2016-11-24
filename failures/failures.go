package failures

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	spew.Dump(err)

	switch err := err.(type) {
	case govalidator.Errors:
		spew.Dump(err)
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "Internal server error",
	})
}
