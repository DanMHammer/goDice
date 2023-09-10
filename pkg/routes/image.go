package routes

import (
	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
	"github.com/gin-gonic/gin"
)

func Image(Cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := Cache.GetResult(id)
		w := c.Writer
		dice.Generate(w, result)
	}
}
