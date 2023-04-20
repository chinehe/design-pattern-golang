package simple

// IntSlice 整型切片
type IntSlice []int

func (s *IntSlice) CreateIterator() *IntSliceIterator {
	return NewIntSliceIterator(*s)
}

// IntSliceIterator 整型切片的迭代器
type IntSliceIterator struct {
	intSlice IntSlice
	index int
}

func NewIntSliceIterator(intSlice []int) *IntSliceIterator {
	return &IntSliceIterator{intSlice: intSlice}
}

func (s *IntSliceIterator)HasNext()  bool{
	if s.index<len(s.intSlice) {
		return true
	}
	return false
}

func (s *IntSliceIterator)Next()  int{
	s.index++
	return s.intSlice[s.index-1]
}

