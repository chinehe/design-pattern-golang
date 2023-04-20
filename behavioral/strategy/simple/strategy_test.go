package simple

import (
	"strconv"
	"testing"
)

func TestStrategy(t *testing.T) {
	cache := NewCache(&FIFOStrategy{}, 10)
	for i := 1; i <= 20; i++ {
		cache.Add(strconv.Itoa(i))
	}
	cache.Print()

	cache2 := NewCache(&LIFOStrategy{}, 10)
	for i := 1; i <= 20; i++ {
		cache2.Add(strconv.Itoa(i))
	}
	cache2.Print()
}
