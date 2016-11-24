package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/constants"
)

func Authentication(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.HeaderToken("apikey", c.Request)
		if err != nil {
			return err
		}

		credentials := strings.SplitN(token, ":", 2)
		if len(credentials) != 2 {
			return ErrInvalidAPIKey
		}

		user, err := managers.GetUserByUsernameAndAPIKey(c, credentials[0], credentials[1])
		if err != nil {
			return ErrInvalidCredentials
		}

		if !user.Active() {
			return ErrUserBlocked
		}

		c.Set(constants.AuthUserKey, user)
		c.Set(constants.AuthAPIKeyKey, credentials[1])
		c.Next()
	}
}
