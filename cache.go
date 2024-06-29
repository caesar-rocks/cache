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

func (cache *Cache) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	set := cache.Client.Set(ctx, key, value, expiration)
	return set.Val(), set.Err()
}

func (cache *Cache) Get(key string) (string, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	get := cache.Client.Get(ctx, key)
	return get.Val(), get.Err()
}

func (cache *Cache) Delete(key string) (int64, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	del := cache.Client.Del(ctx, key)
	return del.Val(), del.Err()
}

func (cache *Cache) FindAllKeys(key string) ([]string, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	var cursor uint64
	var result []string
	for {
		var keys []string
		var err error
		keys, cursor, err = cache.Client.Scan(ctx, cursor, key+"*", 10).Result()
		if err != nil {
			return result, err
		}
		result = append(result, keys...)
		if cursor == 0 {
			break
		}
	}
	return result, nil
}
