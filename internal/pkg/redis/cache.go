package redis

import (
	"dennic_api_gateway/internal/pkg/config"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type RedisDB struct {
	Client *redis.Client
}

func New(cfg *config.Config) (*RedisDB, error) {
	db, err := strconv.Atoi(cfg.Redis.Name)
	if err != nil {
		return &RedisDB{}, err
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       db,
	})
	return &RedisDB{Client: rdb}, nil
}
