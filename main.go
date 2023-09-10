package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/routes"
)

var cacheEngineFlag = flag.String("engine", "gocache", "Storage engine to use for hashes and messages.  Supported: redis, gocache. Default: gocache")

// Cache - Cache Engine for saving results
var Cache cache.CacheEngine

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	Cache, _ = cache.SetupCache(*cacheEngineFlag)

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.Default()

	addRoutes(r)

	r.Run(":" + port)
}

func addRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/roll/:input", routes.Roll(Cache))
	r.GET("/image/:id", routes.Image(Cache))
	r.GET("/rollImage/:input", routes.RollImage())
}
