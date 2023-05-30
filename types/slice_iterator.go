package types

type sliceIterator[T any] struct {
	curr   int
	size   int
	source []T
}

func (i *sliceIterator[T]) Next() T {
	v := i.source[i.curr]
	i.curr++
	return v
}

func (i *sliceIterator[T]) HasNext() bool {
	return i.size > i.curr
}

func NewSliceIterator[T any](src []T) Iterator[T] {
	return &sliceIterator[T]{
		source: src,
		curr:   0,
		size:   len(src),
	}
}

func NewStringIterator(src string) Iterator[rune] {
	return &sliceIterator[rune]{
		source: []rune(src),
		curr:   0,
		size:   len(src),
	}
}
