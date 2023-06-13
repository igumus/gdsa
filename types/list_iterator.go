package types

type listIterator[T comparable] struct {
	current List[T]
}

func NewListIterator[T comparable](coll List[T]) Iterator[T] {
	return &listIterator[T]{
		current: coll,
	}
}

func (i *listIterator[T]) HasNext() bool {
	return i.current != nil && !i.current.IsEmpty()
}

func (i *listIterator[T]) Next() T {
	value := i.current.Get()
	i.current = i.current.Rest()
	return value
}
