package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
	newdice "github.com/danmhammer/goDice/pkg/new-dice"
)

// 4d20H3+3d4L1+12-3
func roll(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		input := c.Param("input")
		result := dice.PerformCommands(input)
		result.Valid = true
		result.Input = input

		id, _ := shortid.Generate()
		baseUrl := os.Getenv("HEROKU_APP_DEFAULT_DOMAIN_NAME")
		result.Image = baseUrl + id

		cache.SaveResult(id, result)
		c.JSON(200, result)
	}
}

func rollJson(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input newdice.RollRequest
		c.BindJSON(&input)
		result, err := newdice.RollDice(input)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// id, _ := shortid.Generate()
		// baseUrl := os.Getenv("HEROKU_APP_DEFAULT_DOMAIN_NAME")
		// result.Image = baseUrl + id

		// cache.SaveResult(id, result)
		c.JSON(200, result)
	}
}
