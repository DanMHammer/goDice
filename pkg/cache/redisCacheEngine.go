package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/danmhammer/goDice/pkg/dice"
	"github.com/danmhammer/goDice/pkg/newdice"
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
func (rdb *RedisEngine) SaveResult(id string, result dice.Result) {

	jsonResult, _ := json.Marshal(result)

	err := rdb.RedisClient.Set(ctx, id, jsonResult, 30*time.Minute).Err()
	if err != nil {
		return
	}

}

// GetResult - Get Result from Redis
func (rdb *RedisEngine) GetResult(id string) dice.Result {
	jsonResult, err := rdb.RedisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		return dice.Result{}
	} else if err != nil {
		panic(err)
	}

	result := dice.Result{}

	if err := json.Unmarshal([]byte(jsonResult), &result); err != nil {
		panic(err)
	}
	return result
}

func (rdb *RedisEngine) SaveRes(id string, res newdice.RollResponse) {
	jsonResult, _ := json.Marshal(res)

	err := rdb.RedisClient.Set(ctx, id, jsonResult, 30*time.Minute).Err()
	if err != nil {
		return
	}
}

func (rdb *RedisEngine) GetRes(id string) newdice.RollResponse {
	jsonResult, err := rdb.RedisClient.Get(ctx, id).Result()
	if err == redis.Nil {
		return newdice.RollResponse{}
	} else if err != nil {
		panic(err)
	}

	result := newdice.RollResponse{}

	if err := json.Unmarshal([]byte(jsonResult), &result); err != nil {
		panic(err)
	}
	return result
}
