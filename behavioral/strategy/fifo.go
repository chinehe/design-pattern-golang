package strategy

// FIFOStrategy 先进先出策略
type FIFOStrategy struct {
}

func (fifo *FIFOStrategy) Remove(cache *Cache) {
	cache.Storage = cache.Storage[1:]
}
