package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/loggers"
	"github.com/juliengroch/todolist/sanitizing"
	"github.com/juliengroch/todolist/store"
)

// ApplicationContextOptions are application options.
type ApplicationContextOptions struct {
	Config    config.Config
	Store     store.Store
	Logger    loggers.Logger
	Sanitizer sanitizing.Sanitizer
}

// ApplicationContext initialize the application context.
func ApplicationContext(opts *ApplicationContextOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, opts.Config)
		store.ToContext(c, opts.Store)
		loggers.ToContext(c, opts.Logger)
		sanitizing.ToContext(c, opts.Sanitizer)

		c.Next()
	}
}
