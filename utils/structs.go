package utils

import (
	"reflect"
	"strings"
)

// ExtractJSONTag return the json tag from a struct without the options
func ExtractJSONTag(o interface{}, field string) string {
	rf, ok := reflect.TypeOf(o).Elem().FieldByName(field)
	if !ok {
		return ""
	}
	tag := rf.Tag.Get("json")
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx]
	}
	return tag
}
