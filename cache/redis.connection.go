package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

type RedisConnectionType struct {
	Client  *redis.Client
	Context context.Context
}

var RedisConnection RedisConnectionType

func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", host, port)
	RedisConnection.Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	RedisConnection.Context = context.Background()
}
