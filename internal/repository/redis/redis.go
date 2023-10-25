package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Addr string
	Pass string
}

var ctx = context.Background()

func InitRedis(cfg *RedisConfig) (*redis.Client, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       0,
	})
	log.Infoln("Check redis...")
	_, err := cache.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
		return nil, err
	}
	return cache, nil
}
