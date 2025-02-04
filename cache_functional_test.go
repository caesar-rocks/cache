package cache

import (
	"fmt"
	"strconv"
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
		t.Fatalf("Did not setup and retrieve properly")
	}
	if res != value {
		t.Fatalf("Did not setup and retrieve properly")
	}
}

func TestDelete(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	cache.Set("keyD", "value", 0)
	res, err := cache.Delete("keyD")
	if err != nil {
		t.Fatalf("Did not delete key")
	}
	if res != 1 {
		t.Fatalf("Did not delete key")
	}
}

func TestFindAllKeys(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	for i := 0; i < 5; i++ {
		cache.Set(fmt.Sprintf("key%d", i), "value"+strconv.Itoa(i), 0)
	}
	res, err := cache.FindAllKeys("key")
	if err != nil {
		t.Fatalf("Error occured when finding keys")
	}
	if len(res) < 4 {
		t.Fatalf("Did not find all keys")
	}
}

func TestFlush(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	for i := 0; i < 5; i++ {
		cache.Set(fmt.Sprintf("flush%d", i), fmt.Sprintf("value%d", i), 0)
	}
	res, _ := cache.FindAllKeys("flush")
	if len(res) < 4 {
		t.Fatalf("Trouble setting up test")
	}
	cache.FlushDB()
	res2, _ := cache.FindAllKeys("flush")
	if len(res2) != 0 {
		t.Fatalf("Didn't flush db")
	}

}

func TestHSetHGet(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	cache.HSet("key", "name", "John")
	res, err := cache.HGet("key", "name")
	if err != nil {
		t.Fatalf("Did not setup and retrieve properly")
	}
	if res != "John" {
		t.Fatalf("Did not setup and retrieve properly")
	}
}

func TestHMSetHMGet(t *testing.T) {
	cache := NewCache(&CacheConfig{
		"redis://default:password@localhost:6379",
		"",
	})
	cache.HMSet("key", map[string]interface{}{
		"name": "John",
		"age":  30,
	})
	res, err := cache.HMGet("key", "name")
	if err != nil {
		t.Fatalf("Did not setup and retrieve properly")
	}
	out := res[0].(string)
	if out != "John" {
		t.Fatalf("Did not setup and retrieve properly")
	}
}
