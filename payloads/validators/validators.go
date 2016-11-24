package validators

import "github.com/asaskevich/govalidator"
import "github.com/juliengroch/todolist/payloads"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func Task(payload *payloads.Task) (bool, error) {
	return govalidator.ValidateStruct(payload)
}
