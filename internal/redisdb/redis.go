package redisdb

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var Client *redis.Client

func InitRedis(addr, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}