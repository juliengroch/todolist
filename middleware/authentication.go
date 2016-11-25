package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/constants"
	"github.com/juliengroch/todolist/managers"
)

// Authentication middleware
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := headerToken("apikey", c.Request)
		if err != nil {
			abordWithErrorToResponse(c, err)
			return
		}

		credentials := strings.SplitN(token, ":", 2)
		if len(credentials) != 2 {
			abordWithErrorToResponse(c, ErrInvalidAPIKey)
			return
		}

		user, errm := managers.GetUserByUsernameAndAPIKey(c, credentials[0], credentials[1])
		if errm != nil {
			abordWithErrorToResponse(c, ErrInvalidCredentials)
			return
		}

		c.Set(constants.AuthUserKey, user)
		c.Set(constants.AuthAPIKeyKey, credentials[1])
		c.Next()
	}
}

/******** HELPER ********/

func abordWithErrorToResponse(c *gin.Context, err *Error) {
	c.JSON(err.Code, gin.H{"message": err.Message})
	c.Abort()
}

// HeaderToken returns authorization header token.
func headerToken(name string, r *http.Request) (string, *Error) {
	name = strings.ToLower(name)

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoHeader
	}

	headerData := strings.Split(authHeader, " ")

	// Remove mutiple spaces
	var authData []string
	for _, str := range headerData {
		if str != "" {
			authData = append(authData, str)
		}
	}

	if len(authData) != 2 {
		return "", ErrInvalidHeaderFormat
	}

	authName := strings.ToLower(strings.TrimSpace(authData[0]))
	if authName != name {
		return "", ErrInvalidHeaderMethod
	}

	return authData[1], nil
}
