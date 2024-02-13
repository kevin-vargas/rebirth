package redis

import (
	"context"
	"time"

	"github.com/kevin-vargas/rebirth/db"
	"github.com/redis/go-redis/v9"
)

type redisDB struct {
	ttl             time.Duration
	currentValueKey string
	masterKey       string
	c               *redis.Client
}

func (r *redisDB) getValue(ctx context.Context, key string) (string, error) {
	val, err := r.c.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", db.ErrNotFound
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisDB) GetMasterIP(ctx context.Context) (string, error) {
	return r.getValue(ctx, r.masterKey)
}

func (r *redisDB) GetCurrentIP(ctx context.Context) (string, error) {
	return r.getValue(ctx, r.currentValueKey)
}

func (r *redisDB) SetCurrentIP(ctx context.Context, ip string) error {
	err := r.c.Set(ctx, r.currentValueKey, ip, r.ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func New(a, p, mk string, cvk string) db.DB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     a,
		Password: p,
		DB:       0,
	})

	return &redisDB{
		currentValueKey: cvk,
		masterKey:       mk,
		c:               rdb,
	}
}
