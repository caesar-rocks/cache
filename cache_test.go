package cache

import (
	"testing"
)

// //////////////////////////////////////////////////// //
// Tests are NOT mocked, just for development purposes  //
// //////////////////////////////////////////////////// //
func TestSetGet(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	value := "value"
	cache.Set("key", value, 0)
	res, err := cache.Get("key")
	if err != nil {
		panic(err)
	}
	if res != value {
		t.Fatalf("Did not setup and retreive properly")
	}
}
