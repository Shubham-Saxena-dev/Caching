package cache

type EvictionStrategy interface {
	Evict(*cacheClient)
	Type() string
}
