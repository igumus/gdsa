package types

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

type Collection[T any] interface {
	IsEmpty() bool
	Count() int
	ContainsValue(T) bool
}

type List[T any] interface {
	Collection[T]
	Get() T
	Rest() List[T]
	Add(T) List[T]
	Equals(List[T]) bool
}

type Indexed[T any] interface {
	Collection[T]
	Set(int, T)
	Get(int) (T, bool)
}
