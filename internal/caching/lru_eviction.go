package cache

const lru = "LRU"

type lruEviction struct{}

func NewLruEviction() EvictionStrategy {
	return &lruEviction{}
}

func (l *lruEviction) Evict(_ *cacheClient) {
	// lru eviction logic
}

func (l *lruEviction) Type() string {
	return lru
}
