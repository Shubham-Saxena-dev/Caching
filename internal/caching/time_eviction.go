package cache

import "time"

const TimeEviction = "time"

type timeEviction struct{}

func NewTimeEviction() EvictionStrategy {
	return &timeEviction{}
}

func (t *timeEviction) Evict(c *cacheClient) {
	var oldKey string
	var oldTime time.Time

	for key, item := range c.items {
		if oldTime.IsZero() || item.expiration.Before(oldTime) { // we should evict the oldest only.
			oldTime = item.expiration
			oldKey = key
		}
	}
	delete(c.items, oldKey)
}

func (t *timeEviction) Type() string {
	return TimeEviction
}
