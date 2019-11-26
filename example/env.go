package example

import (
	"github.com/go-redis/redis"
)

type RedisEnv struct{}

func (*RedisEnv) Setup() {
	client := redis.NewClient(&redis.Options{DB: 10})
	client.Set("users:100:nickname", "me", 0)
}

func (*RedisEnv) TearDown() {
	client := redis.NewClient(&redis.Options{DB: 10})
	client.FlushDB()
}
