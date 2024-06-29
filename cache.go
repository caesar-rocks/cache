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

func (cache *Cache) FlushDB() error {
	ctx := context.Background()
	return cache.Client.FlushDB(ctx).Err()
}

func (cache *Cache) HSet(key string, setKey string, setValue interface{}) (int64, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	set := cache.Client.HSet(ctx, key, setKey, setValue)
	return set.Val(), set.Err()
}

func (cache *Cache) HGet(key string, setKey string) (string, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	get := cache.Client.HGet(ctx, key, setKey)
	return get.Val(), get.Err()
}

func (cache *Cache) HMSet(key string, setValue map[string]interface{}) (bool, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	set := cache.Client.HMSet(ctx, key, setValue)
	return set.Val(), set.Err()
}

func (cache *Cache) HMGet(key string, setKey string) ([]interface{}, error) {
	ctx := context.Background()
	key = cache.Prefix + key
	set := cache.Client.HMGet(ctx, key, setKey)
	return set.Val(), set.Err()
}
