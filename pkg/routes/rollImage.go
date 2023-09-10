package routes

import (
	"github.com/danmhammer/goDice/pkg/dice"
	"github.com/gin-gonic/gin"
)

func RollImage() func(c *gin.Context) {
	return func(c *gin.Context) {
		input := c.Param("input")
		result := dice.PerformCommands(input)
		result.Input = input
		w := c.Writer
		dice.Generate(w, result)
	}
}
