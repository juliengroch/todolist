package failures

import (
	"errors"
	"net/http"

	"github.com/mholt/binding"

	"github.com/juliengroch/todolist/constants"
)

// Global errors
var (
	ErrWrongStartCmdCli = errors.New("Wrong command to use the app, use run or migrate")
)

// HTTPError struct for http errors
type HTTPError struct {
	HTTPStatus int            `json:"-"`
	Message    interface{}    `json:"message"`
	Type       string         `json:"type,omitempty"`
	Err        error          `json:"-"`
	Errors     binding.Errors `json:"errors,omitempty"`
}

func (e HTTPError) Error() string {
	return e.Message.(string)
}

// ValidationError formats govalidator errors
func ValidationError(errs binding.Errors) HTTPError {
	code := 422

	if len(errs) > 0 && errs[0].Classification == "ContentTypeError" {
		code = http.StatusUnsupportedMediaType
	}

	return HTTPError{
		HTTPStatus: code,
		Message:    constants.ValidationFailed,
		Errors:     errs,
		Type:       "validation_failed",
	}
}
