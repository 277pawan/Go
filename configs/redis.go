package configs

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx     = context.Background()
	RedisDB *redis.Client
)

func ConnectRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisDB.Ping(Ctx).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println("Redis Connected")
}
