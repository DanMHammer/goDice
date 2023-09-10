package cache

import (
	"context"
	"errors"

	"github.com/danmhammer/goDice/pkg/dice"
)

// CacheEngine interface
type CacheEngine interface {
	SaveRes(id string, res dice.RollResponse)
	GetRes(id string) (res dice.RollResponse)
}

var ctx = context.Background()

// SetupCache - Create Cache Engine
func SetupCache(cacheEngineFlag string) (output CacheEngine, err error) {
	switch cacheEngine := cacheEngineFlag; cacheEngine {
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
