package simple

// RemoveStrategy 移除策略
type RemoveStrategy interface {
	Remove(cache *Cache)
}
