package simple

// LIFOStrategy 后进先出策略
type LIFOStrategy struct {
	
}

func (lifo *LIFOStrategy) Remove(cache *Cache) {
	cache.Storage = cache.Storage[:len(cache.Storage)-1]
}



