package types

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

type Collection[T any] interface {
	IsEmpty() bool
	Count() int
	ContainsValue(T) bool
	Add(T) Collection[T]
}

type List[T any] interface {
	Collection[T]
	Get() T
	Rest() List[T]
}

type Indexed[T any] interface {
	Collection[T]
	Set(int, T)
	Get(int) (T, bool)
}
