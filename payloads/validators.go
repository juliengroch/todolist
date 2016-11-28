package payloads

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/mholt/binding"

	"github.com/juliengroch/todolist/utils"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Validate payload
func Validate(req *http.Request, errs binding.Errors, data interface{}) binding.Errors {
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

	return ValidateBinding(errs, data)
}

// ValidateBinding payload with Binding
func ValidateBinding(errs binding.Errors, data interface{}, prefix ...string) binding.Errors {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		fieldErrs := govalidator.ErrorsByField(err)

		var fieldName string
		for field, msg := range fieldErrs {
			fieldName = utils.ExtractJSONTag(data, field)
			if fieldName == "" {
				continue
			}

			if len(prefix) > 0 {
				fieldName = fmt.Sprintf("%s.%s", prefix[0], fieldName)
			}

			if msg == "non zero value required" {
				errs = append(errs, binding.Error{
					FieldNames:     []string{fieldName},
					Classification: "RequiredError",
					Message:        "Required field",
				})
			} else {
				errs = append(errs, binding.Error{
					FieldNames:     []string{fieldName},
					Classification: "InvalidValueError",
					Message:        "Invalid field",
				})
			}
		}
	}
	return errs
}
