package payloads

import (
	"net/http"

	"github.com/juliengroch/todolist/models"
	"github.com/mholt/binding"
)

// Comment payload for POST (create)
type Comment struct {
	Message string       `json:"message" valid:"required"`
	TaskID  string       `json:"taskid" valid:"required"`
	User    *models.User `json:"-" valid:"optional"`
}

// FieldMap for payload (github.com/mholt/binding)
func (t *Comment) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{}
}

// Validate payload
func (t *Comment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
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

// CommentUp payload for PATCH (update)
type CommentUp struct {
	Message string `json:"message" valid:"required"`
}

// FieldMap for payload (github.com/mholt/binding)
func (t *CommentUp) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{}
}

// Validate payload
func (t *CommentUp) Validate(req *http.Request, errs binding.Errors) binding.Errors {
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
