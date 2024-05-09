package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	slice := IntSlice{1, 2, 3, 4, 5, 6, 7, 8, 9}
	iterator := slice.CreateIterator()
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
