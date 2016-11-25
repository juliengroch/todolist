package application

import (
	"context"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/store"
)

// Load application
func Load(cfg *config.Config) (context.Context, error) {

	// Empty context
	ctx := context.Background()

	s, err := store.New(cfg.Database)

	if err != nil {
		return nil, err
	}

	// Add config to context
	ctx = config.NewContext(ctx, *cfg)

	// Add store to context
	ctx = store.NewContext(ctx, s)

	// Add logger to context
	ctx = loggers.NewContext(ctx, loggers.New())

	return ctx, nil
}

// Migrate database fonction for CLI
func Migrate(ctx context.Context) error {
	return store.Migrate(ctx)
}
