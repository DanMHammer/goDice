package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func saveJSON(id string, result Result) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	jsonResult, _ := json.Marshal(result)

	err := rdb.Set(ctx, id, jsonResult, 15*time.Minute).Err()
	if err != nil {
		panic(err)
	}

}

func getJSON(id string) string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	jsonResult, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return jsonResult
	}

}
