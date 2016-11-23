package config

import (
	"context"
)

// ContextKey is the context key.
const ContextKey = "configuration"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Config associated with this context.
func FromContext(c context.Context) Config {
	return c.Value(ContextKey).(Config)
}

// ToContext adds the Config to this context if it supports
// the Setter interface.
func ToContext(c Setter, s Config) {
	c.Set(ContextKey, s)
}

// NewContext returns a new context with the given config.
func NewContext(ctx context.Context, c Config) context.Context {
	return context.WithValue(ctx, ContextKey, c)
}
