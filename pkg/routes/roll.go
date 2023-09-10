package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/dice"
)

func storeResult(cache cache.CacheEngine, result *dice.RollResponse) {
	id, _ := shortid.Generate()
	baseUrl := os.Getenv("HEROKU_APP_DEFAULT_DOMAIN_NAME")
	result.ImageUrl = fmt.Sprintf("%s/image/%s", baseUrl, id)

	cache.SaveRes(id, *result)
}

// 4d20H3+3d4L1
// roll
// @Summary roll dice based on string input
// @Tags roll
// @Produce json
// @Param input path string true "input"
// @Success 200 object dice.RollResponse
// @Router /roll/{input} [get]
func roll(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		input := c.Param("input")
		req, err := dice.ParseStringInput(input)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := dice.RollDice(req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		storeResult(cache, &result)

		c.JSON(200, result)
	}
}

// rollJSON
// @Summary roll dice based on json post
// @Tags roll
// @Accept json
// @Produce json
// @Param roll body dice.RollRequest true "roll"
// @Success 200 object dice.RollResponse
// @Router /roll [post]
func rollJson(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		var input dice.RollRequest
		c.BindJSON(&input)

		result, err := dice.RollDice(input)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		storeResult(cache, &result)

		c.JSON(200, result)
	}
}

// rollImage
// @Summary roll dice based on string input and return image
// @Tags roll
// @Param input path string true "input"
// @Router /rollImage/{input} [get]
func rollImage(cache cache.CacheEngine) func(c *gin.Context) {
	return func(c *gin.Context) {
		input := c.Param("input")
		req, err := dice.ParseStringInput(input)

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := dice.RollDice(req)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		storeResult(cache, &result)

		c.Redirect(http.StatusFound, result.ImageUrl)
	}
}
