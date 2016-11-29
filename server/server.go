package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/middleware"
	"github.com/juliengroch/todolist/sanitizing"
	"github.com/juliengroch/todolist/store"
	"github.com/juliengroch/todolist/views"
)

// Run fonction to start the server
func Run(ctx context.Context) error {
	router := Router(ctx)

	views.Routes(router)

	router.Run(fmt.Sprintf(":%d", config.FromContext(ctx).Server.Port))

	return nil
}

// Router init a gin router with ctx
func Router(ctx context.Context) *gin.Engine {
	router := gin.Default()

	appContext := &middleware.ApplicationContextOptions{
		Config:    config.FromContext(ctx),
		Store:     store.FromContext(ctx),
		Logger:    loggers.FromContext(ctx),
		Sanitizer: sanitizing.FromContext(ctx),
	}

	router.Use(middleware.ApplicationContext(appContext))

	return router
}
