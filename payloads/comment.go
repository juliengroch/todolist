package payloads

import (
	"net/http"

	"github.com/juliengroch/todolist/models"
	"github.com/mholt/binding"
)

// Comment payload for POST (create)
type Comment struct {
	Message string       `json:"message" valid:"required"`
	TaskID  string       `json:"-" valid:"optional"`
	User    *models.User `json:"-" valid:"optional"`
}

// FieldMap for payload (github.com/mholt/binding)
func (c *Comment) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{}
}

// Validate Comment payload
func (c *Comment) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	return Validate(req, errs, c)
}
