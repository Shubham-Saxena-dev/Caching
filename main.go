package main

import (
	"fmt"
	cache2 "semrush/internal/caching"
	"time"
)

func main() {

	fmt.Println("Hi, this is semrush exercise")

	cache := cache2.GetCache(3, 10*time.Second, cache2.NewTimeEviction())
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	fmt.Println(cache.CacheSize())
	if value, err := cache.Get("key1"); err == nil {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 ", err.Error())
	}

	time.Sleep(6 * time.Second)

	cache.Set("key3", "value3")

	if value, err := cache.Get("key2"); err == nil {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 not found")
	}

	if value, err := cache.Get("key3"); err != nil {
		fmt.Println("key3:", value)
	} else {
		fmt.Println("key3 not found")
	}
}
