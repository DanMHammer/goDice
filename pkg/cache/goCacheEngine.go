package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"

	"github.com/danmhammer/goDice/pkg/dice"
)

// GoCacheEngine structure
type GoCacheEngine struct {
	Cache *gocache.Cache
}

// Connect - Create Go Cache
func (gc *GoCacheEngine) Connect() (err error) {
	gc.Cache = gocache.New(5*time.Minute, 10*time.Minute)
	return
}

// NewGoCacheEngine - Instantiate GoCache
func NewGoCacheEngine() (output *GoCacheEngine, err error) {
	var engine GoCacheEngine
	err = engine.Connect()
	if err != nil {
		return
	}
	return &engine, nil
}

func (gc *GoCacheEngine) SaveRes(id string, res dice.RollResponse) {
	gc.Cache.Set(id, res, gocache.DefaultExpiration)
}

func (gc *GoCacheEngine) GetRes(id string) dice.RollResponse {
	if x, found := gc.Cache.Get(id); found {
		result := x.(dice.RollResponse)
		return result
	}
	return dice.RollResponse{}
}
