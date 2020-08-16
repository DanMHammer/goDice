package main

import (
	"time"

	"github.com/patrickmn/go-cache"
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
func (gc *GoCacheEngine) SaveResult(id string, result Result) {
	gc.Cache.Set(id, result, cache.DefaultExpiration)
}

// GetResult - Get Result from Cache
func (gc *GoCacheEngine) GetResult(id string) Result {
	if x, found := gc.Cache.Get(id); found {
		result := x.(Result)
		return result
	}
	return Result{}
}
