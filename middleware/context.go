package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juliengroch/todolist/config"
	"github.com/juliengroch/todolist/store"
)

func SetConfig(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		config.ToContext(c, cfg)
		c.Next()
	}
}

func Config(c *gin.Context) config.Config {
	return c.MustGet("config").(config.Config)
}

func SetStore(b store.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		store.ToContext(c, b)
		c.Next()
	}
}

func Store(c *gin.Context) store.Store {
	return c.MustGet("store").(store.Store)
}
