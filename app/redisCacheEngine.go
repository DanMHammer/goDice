package main

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisEngine structure
type RedisEngine struct {
	RedisClient *redis.Client
}

// Connect - Connect to Redis
func (rdb *RedisEngine) Connect() (err error) {
	rdb.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.RedisClient.Ping(ctx).Result()

	return
}

// NewRedisEngine - Instantiate Redis
func NewRedisEngine() (output *RedisEngine, err error) {
	var engine RedisEngine
	err = engine.Connect()
	if err != nil {
		return
	}
	return &engine, nil
}

// SaveResult - Save Result to Redis
func (rdb *RedisEngine) SaveResult(id string, result Result) {

	jsonResult, _ := json.Marshal(result)

	err := rdb.RedisClient.Set(ctx, id, jsonResult, 30*time.Minute).Err()
	if err != nil {
		return
	}

}

// GetResult - Get Result from Redis
func (rdb *RedisEngine) GetResult(id string) Result {

	jsonResult, err := rdb.RedisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		return Result{}
	} else if err != nil {
		panic(err)
	}

	result := Result{}

	if err := json.Unmarshal([]byte(jsonResult), &result); err != nil {
		panic(err)
	}
	return result
}
