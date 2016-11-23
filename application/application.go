package application

import (
	"context"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/store"
)

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

	return ctx, nil
}
