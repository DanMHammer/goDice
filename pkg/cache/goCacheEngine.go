package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/danmhammer/goDice/pkg/dice"
)

// GoCacheEngine structure
type GoCacheEngine struct {
	Cache *cache.Cache
}

// Connect - Create Go Cache
func (gc *GoCacheEngine) Connect() (err error) {
	gc.Cache = cache.New(5*time.Minute, 10*time.Minute)
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

// SaveResult - Save Result to Cache
func (gc *GoCacheEngine) SaveResult(id string, result dice.Result) {
	gc.Cache.Set(id, result, cache.DefaultExpiration)
}

// GetResult - Get Result from Cache
func (gc *GoCacheEngine) GetResult(id string) dice.Result {
	if x, found := gc.Cache.Get(id); found {
		result := x.(dice.Result)
		return result
	}
	return dice.Result{}
}
