package routes

import (
	"net/http"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/danmhammer/goDice/docs"
)

// @BasePath /
func AddRoutes(r *gin.Engine, c cache.CacheEngine) {
	docs.SwaggerInfo.BasePath = "/"

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/roll/:input", roll(c))
	r.POST("/roll", rollJson(c))
	r.GET("/image/:id", image(c))
	r.GET("/image-new/:id", imageNew(c))
	r.GET("/rollImage/:input", rollImage())
}
