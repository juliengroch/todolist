package sanitizing

import (
	"context"
)

// ContextKey is the context key.
const ContextKey = "sanitizer"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Sanitizer instance associated with this context.
func FromContext(c context.Context) Sanitizer {
	return c.Value(ContextKey).(Sanitizer)
}

// ToContext adds the Sanitizer to this context if it supports the Setter interface.
func ToContext(c Setter, s Sanitizer) {
	c.Set(ContextKey, s)
}

// NewContext returns a new context with sanitizer instance.
func NewContext(ctx context.Context, s Sanitizer) context.Context {
	return context.WithValue(ctx, ContextKey, s)
}
