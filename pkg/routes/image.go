package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
)

func image(Cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := Cache.GetRes(id)
		w := c.Writer
		dice.Render(w, result)
	}
}
