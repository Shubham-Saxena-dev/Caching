package cache

import (
	"semrush/internal/customerrors"
	"sync"
	"time"
)

//const DefaultExpiration time.Duration = 10 * time.Second

var (
	once  sync.Once
	cache Cache
)

type Cache interface {
	Set(string, interface{})
	Get(string) (interface{}, error)
	Remove(string)
	Evict()
	EvictExpired()
	CacheSize() int
}

type cacheClient struct {
	capacity         int
	items            map[string]cacheItem
	expireAfter      time.Duration
	mutex            sync.Mutex
	evictionStrategy EvictionStrategy
}

func GetCache(capacity int, expiration time.Duration, strategy EvictionStrategy) Cache {
	once.Do(func() {
		cache = &cacheClient{
			capacity:         capacity,
			items:            make(map[string]cacheItem),
			expireAfter:      expiration,
			evictionStrategy: strategy,
		}
	})
	if strategy.Type() == TimeEviction {
		go cache.EvictExpired()
	}
	return cache
}

func (c *cacheClient) EvictExpired() {
	for {
		time.Sleep(c.expireAfter / 2)
		c.mutex.Lock()
		for key, item := range c.items {
			if item.IsExpired() {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *cacheClient) Get(key string) (interface{}, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, ok := c.items[key]
	if !ok {
		return nil, customerrors.ErrorInvalidKey(key, "not found")
	}

	item.expiration = time.Now().Add(c.expireAfter)
	return item.value, nil
}

func (c *cacheClient) Set(key string, val interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.items) >= c.capacity {
		c.Evict()
	}

	c.items[key] = cacheItem{
		key:        key,
		value:      val,
		expiration: time.Now().Add(c.expireAfter),
	}
}

func (c *cacheClient) Evict() {
	c.evictionStrategy.Evict(c)
}

func (c *cacheClient) Remove(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.items[key]; !ok {
		return
	}
	delete(c.items, key)
}

func (c *cacheClient) CacheSize() int {
	return len(c.items)
}
