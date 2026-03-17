package redis

import (
	"context"
	"errors"
	"fmt"
	"gateway/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisRepo struct {
	client *redis.Client
}

func NewRedisClient() (*RedisRepo, error) {
	cfg := config.Load()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_PORT,
		Password: cfg.REDIS_PASS,
		DB:       cfg.REDIS_DB,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("redis connection failed: %w", err)
	}

	return &RedisRepo{client: rdb}, nil
}

func (r *RedisRepo) BlacklistToken(token string, exp time.Duration) error {
	if token == "" {
		return errors.New("token cannot be empty")
	}
	return r.client.Set(ctx, "blacklist:"+token, "blacklisted", exp).Err()
}

func (r *RedisRepo) IsBlacklisted(token string) (bool, error) {
	if token == "" {
		return false, errors.New("token cannot be empty")
	}

	val, err := r.client.Get(ctx, "blacklist:"+token).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("redis get error: %w", err)
	}

	return val == "blacklisted", nil
}

func (r *RedisRepo) Close() error {
	return r.client.Close()
}
