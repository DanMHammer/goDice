package routes

import (
	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
	"github.com/danmhammer/goDice/pkg/newdice"
	"github.com/gin-gonic/gin"
)

func image(Cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := Cache.GetResult(id)
		w := c.Writer
		dice.Generate(w, result)
	}
}

func imageNew(Cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := Cache.GetRes(id)
		w := c.Writer
		newdice.GenerateImage(w, result)
	}
}
