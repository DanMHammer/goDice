package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
)

// 4d20H3+3d4L1+12-3
func Roll(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		input := c.Param("input")
		result := dice.PerformCommands(input)
		result.Valid = true
		result.Input = input

		id, _ := shortid.Generate()
		result.Image = "https://k8s.danhammer.dev/image/" + id

		cache.SaveResult(id, result)
		c.JSON(200, result)
	}
}
