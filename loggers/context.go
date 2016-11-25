package loggers

import (
	"context"
)

// ContextKey is the context key.
const ContextKey = "logger"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Logger associated with this context.
func FromContext(c context.Context) Logger {
	return c.Value(ContextKey).(Logger)
}

// ToContext adds the Logger to this context if it supports
// the Setter interface.
func ToContext(c Setter, logger Logger) {
	c.Set(ContextKey, logger)
}

// NewContext creates a new context and adds the Logger
func NewContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ContextKey, logger)
}
