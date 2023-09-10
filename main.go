package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/danmhammer/goDice/pkg/cache"
	"github.com/danmhammer/goDice/pkg/routes"
)

var cacheEngineFlag = flag.String("engine", "gocache", "Storage engine to use for hashes and messages.  Supported: redis, gocache. Default: gocache")

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	c, err := cache.SetupCache(*cacheEngineFlag)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.Default()

	routes.AddRoutes(r, c)

	r.Run(":" + port)
}
