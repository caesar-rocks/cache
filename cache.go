package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheConfig struct {
	URL    string
	Prefix string
}

type Cache struct {
	Client *redis.Client
	Prefix string
}

func NewCache(cfg *CacheConfig) *Cache {
	var db *redis.Client
	opts, err := redis.ParseURL(cfg.URL)
	if err != nil {
		log.Fatal(err)
	}
	db = redis.NewClient(opts)
	return &Cache{db, cfg.Prefix}
}

func (cache *Cache) Set(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	key = cache.Prefix + key
	err := cache.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cache *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	val, err := cache.Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
