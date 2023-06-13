package transducer

import (
	"fmt"

	"github.com/igumus/gdsa/types"
)

type Function[T, R any] func(T) R
type Predicate[T any] Function[T, bool]
type ReducerFunction[T any] func(any, *bool, ...T) any
type Transducer[A, B any] func(ReducerFunction[A]) ReducerFunction[B]

func Append[A any](output any, reduced *bool, items ...A) any {
	if len(items) == 0 || *reduced {
		return output
	}

	item := items[0]
	switch acc := output.(type) {
	case []A:
		acc = append(acc, item)
		return acc
	case types.List[A]:
		acc = acc.Add(item)
		return acc
	default:
		fmt.Printf("unknown type for append: %T\n", acc)
		return acc

	}
}

// TODO (@igumus): add Transduce function
// TODO (@igumus): add several reduce function on last parameter.

func Reduce[A, T any](rfn ReducerFunction[T], initial A, it types.Iterator[T]) A {
	reduced := false
	acc := initial
	for it.HasNext() && !reduced {
		acc = rfn(acc, &reduced, it.Next()).(A)
	}
	return rfn(acc, &reduced).(A)
}

func Combine[A, B, C any](left Transducer[B, C], right Transducer[A, B]) Transducer[A, C] {
	return func(rf ReducerFunction[A]) ReducerFunction[C] {
		return left(right(rf))
	}
}

// Creates a transducer that transforms a reducing function by applying a mapping
// function to each input.
func Map[A, B any](f Function[B, A]) Transducer[A, B] {
	return func(rf ReducerFunction[A]) ReducerFunction[B] {
		return func(acc any, reduced *bool, items ...B) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				value := f(item)
				acc = rf(acc, reduced, value)
			}
			// completing phase
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function by applying a
// predicate to each input and processing only those inputs for which the
// predicate is true.
func Filter[A any](f Predicate[A]) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				if f(item) {
					acc = rf(acc, reduced, item)
				}
			}
			// completing phase
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function such that
// it only processes n inputs, then the reducing process stops.
func Take[A any](n int) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		taken := 0
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				if taken < n {
					acc = rf(acc, reduced, item)
					taken += 1
				} else {
					*reduced = true
				}
			} else {
				// completing phase
				taken = 0
			}
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function such that
// it processes inputs as long as the provided predicate returns true.
// If the predicate returns false, the reducing process stops.
func TakeWhile[A any](f Predicate[A]) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				ret := f(item)
				if ret {
					acc = rf(acc, reduced, item)
				}
				*reduced = !ret
			}
			// completing phase
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function such that
// it skips n inputs, then processes the rest of the inputs.
func Skip[A any](n int) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		skipped := 0
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				if skipped < n {
					skipped += 1
				} else {
					acc = rf(acc, reduced, item)
				}
			} else {
				// completing phase
				skipped = 0
			}
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function such that
// it skips inputs as long as the provided predicate returns true.
// Once the predicate returns false, the rest of the inputs are
// processed.
func SkipWhile[A any](f Predicate[A]) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		drop := true
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				if drop && f(item) {
					return acc
				}
				drop = false
				acc = rf(acc, reduced, item)
			} else {
				// completing phase
				drop = true
			}
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function such that
// it processes every nth input.
func EveryNth[A any](n int) Transducer[A, A] {
	return func(rf ReducerFunction[A]) ReducerFunction[A] {
		nth := 0
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				if nth%n == 0 {
					item := items[0]
					acc = rf(acc, reduced, item)
				}
				nth += 1
			} else {
				// completing phase
				nth = 0
			}
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function that processes
// iterables of input into a reducing function that processes individual inputs
// by gathering series of inputs into partitions of a given size, only forwarding
// them to the next reducing function when enough inputs have been accrued. Processes
// any remaining buffered inputs when the reducing process completes.
func PartitionAll[A any](n int) Transducer[types.Iterator[A], A] {
	return func(rf ReducerFunction[types.Iterator[A]]) ReducerFunction[A] {
		var part []A = nil
		amount := 0
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				if part == nil {
					part = make([]A, 0)
				}
				item := items[0]
				if amount == n {
					acc = rf(acc, reduced, types.NewSliceIterator(part))
					part = make([]A, 0)
					amount = 0
				}
				part = append(part, item)
				amount += 1
			} else {
				// completing phase
				if len(part) > 0 {
					acc = rf(acc, reduced, types.NewSliceIterator(part))
					part = nil
				}
				amount = 0
			}
			return acc
		}
	}
}

// Creates a transducer that transforms a reducing function that processes
// iterables of input into a reducing function that processes individual inputs
// by gathering series of inputs for which the provided partitioning function returns
// the same value, only forwarding them to the next reducing function when the value
// the partitioning function returns for a given input is different from the value
// returned for the previous input.
func PartitionBy[A, B comparable](f Function[A, B]) Transducer[types.Iterator[A], A] {
	return func(rf ReducerFunction[types.Iterator[A]]) ReducerFunction[A] {
		part := make([]A, 0)
		initialized := false
		var prior B
		return func(acc any, reduced *bool, items ...A) any {
			if len(items) > 0 && !*reduced {
				// stepping phase
				item := items[0]
				val := f(item)
				if !initialized || prior == val {
					initialized = true
				} else {
					acc = rf(acc, reduced, types.NewSliceIterator(part))
					part = make([]A, 0)
				}
				prior = val
				part = append(part, item)
			} else {
				// completing phase
				if len(part) > 0 {
					acc = rf(acc, reduced, types.NewSliceIterator(part))
				}
				part = nil
			}
			return acc
		}
	}
}
