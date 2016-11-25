package middleware

import (
	"fmt"
	"net/http"
)

// Error for middleware
type Error struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

// Errors
var (
	ErrInvalidAPIKey       = &Error{Code: http.StatusUnauthorized, Message: "Invalid API Key"}
	ErrInvalidCredentials  = &Error{Code: http.StatusUnauthorized, Message: "Invalid credentials"}
	ErrNoHeader            = &Error{Code: http.StatusUnauthorized, Message: "no header Authorization"}
	ErrInvalidHeaderFormat = &Error{Code: http.StatusUnauthorized, Message: `Authorization: invalid header format, do 'apikey username:your_api_key'`}
	ErrInvalidHeaderMethod = &Error{Code: http.StatusUnauthorized, Message: `Authorization: invalid method, use 'apikey'`}
)
