package cache

import "time"

type cacheItem struct {
	key        string
	value      interface{}
	expiration time.Time
}

func (c *cacheItem) IsExpired() bool {
	return time.Now().After(c.expiration)
}
