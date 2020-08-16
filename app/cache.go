package main

import (
	"context"
	"errors"
)

// CacheEngine interface
type CacheEngine interface {
	// Saves the result to cache
	SaveResult(id string, result Result)
	// Retrieves result from cache
	GetResult(id string) (result Result)
}

var ctx = context.Background()

// SetupCache - Create Cache Engine
func SetupCache() (output CacheEngine, err error) {
	switch cacheEngine := *cacheEngineFlag; cacheEngine {

	case "redis":
		output, err = NewRedisEngine()
		return
	case "gocache":
		output, err = NewGoCacheEngine()
		return
	default:
		err = errors.New("cache engine not supported" + cacheEngine)
		return
	}
}
