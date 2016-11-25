package payloads

import (
	"net/http"

	"github.com/juliengroch/todolist/models"
	"github.com/mholt/binding"
)

// Task payload for POST (create) and  PATCH (update)
type Task struct {
	Title       string       `json:"title" valid:"required,stringlength(4,20)"`
	Description string       `json:"description" valid:"required"`
	Priority    int8         `json:"priority" valid:"required"`
	User        *models.User `json:"-" valid:"optional"`
}

// FieldMap for payload (github.com/mholt/binding)
func (t *Task) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{}
}

// Validate payload
func (t *Task) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	var errFields []string

	if len(errFields) > 0 {
		for _, err := range errFields {
			errs = append(errs, binding.Error{
				FieldNames:     []string{err},
				Classification: "InvalidValueError",
				Message:        "Invalid field",
			})
		}
	}

	return ValidateBinding(errs, t)
}
