package store

import (
	"context"
)

// ContextKey is the context key.
const ContextKey = "store"

// Setter defines a context that enables setting values.
type Setter interface {
	Set(string, interface{})
}

// FromContext returns the Store associated with this context.
func FromContext(c context.Context) Store {
	return c.Value(ContextKey).(Store)
}

// ToContext adds the Store to this context if it supports
// the Setter interface.
func ToContext(c Setter, store Store) {
	c.Set(ContextKey, store)
}

// NewContext creates a new context and adds the store
func NewContext(ctx context.Context, store Store) context.Context {
	return context.WithValue(ctx, ContextKey, store)
}
