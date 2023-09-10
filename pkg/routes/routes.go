package routes

import (
	"net/http"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine, c cache.CacheEngine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/roll/:input", roll(c))
	r.POST("/roll", rollJson(c))
	r.GET("/image/:id", image(c))
	r.GET("/rollImage/:input", rollImage())
}
