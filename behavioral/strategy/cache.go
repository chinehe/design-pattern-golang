package strategy

import "fmt"

// Cache 缓存
type Cache struct {
	Storage        []string
	RemoveStrategy RemoveStrategy
	MaxCapacity    int
}

func NewCache(removeStrategy RemoveStrategy, maxCapacity int) *Cache {
	return &Cache{
		Storage:        []string{},
		RemoveStrategy: removeStrategy,
		MaxCapacity:    maxCapacity,
	}
}

func (cache *Cache) Add(value string) {
	if len(cache.Storage) >= cache.MaxCapacity {
		cache.RemoveStrategy.Remove(cache)
	}
	cache.Storage = append(cache.Storage, value)
}

func (cache *Cache) Print() {
	fmt.Printf("cache:len=%v,values:%v\n", len(cache.Storage), cache.Storage)
}
