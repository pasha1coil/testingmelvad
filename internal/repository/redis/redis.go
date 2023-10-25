package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedisConfig struct {
	Addr string
	Port string
	Pass string
}

func InitRedis(cfg *RedisConfig) (*redis.Client, error) {
	cache := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr + ":" + cfg.Port,
		Password: cfg.Pass,
		DB:       0,
	})

	log.Infoln("Check redis...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := cache.Ping(ctx).Result()
	if err != nil {
		log.Errorf("Failed to connect to Redis:%s", err)
		return nil, err
	}
	return cache, nil
}
