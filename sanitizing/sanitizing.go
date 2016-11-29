package sanitizing

import (
	"github.com/microcosm-cc/bluemonday"
)

// Sanitizer is the interface to sanitize data.
type Sanitizer interface {
	Sanitize(input string) string
}

type bm struct {
	policy *bluemonday.Policy
}

// Sanitize implements Sanitizer Sanitize() method.
func (b bm) Sanitize(input string) string {
	return b.policy.Sanitize(input)
}

// NewBluemonday returns a Bluemonday sanitizer instance.
func NewBluemonday() Sanitizer {
	return bm{
		policy: bluemonday.UGCPolicy(),
	}
}
