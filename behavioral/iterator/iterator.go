package iterator

// IntIterator 整型迭代器接口
type IntIterator interface {
	HasNext() bool
	Next() int
}
