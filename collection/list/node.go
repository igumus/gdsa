package list

import (
	"github.com/igumus/gdsa/types"
)

type node[T comparable] struct {
	count int
	value T
	next  types.List[T]
}

func emptyNode[T comparable]() *node[T] {
	return &node[T]{
		count: 0,
		next:  nil,
	}
}

func newNode[T comparable](v T) *node[T] {
	return &node[T]{
		count: 1,
		value: v,
		next:  nil,
	}
}

func newListFromArray[T comparable](items []T) *node[T] {
	var root *node[T] = emptyNode[T]()
	count := len(items)
	if count > 0 {
		last := count - 1
		root = newNode(items[last])
		for i := last - 1; i >= 0; i-- {
			root = &node[T]{
				count: root.count + 1,
				value: items[i],
				next:  root,
			}
		}
	}
	return root
}

func (n *node[T]) Count() int {
	if n == nil {
		return 0
	}
	return n.count
}

func (n *node[T]) IsEmpty() bool {
	return n == nil || n.Count() == 0
}

func (n *node[T]) Add(val T) types.List[T] {
	if n == nil || n.IsEmpty() {
		return newNode(val)
	}
	return &node[T]{
		count: n.count + 1,
		value: val,
		next:  n,
	}
}

func (n *node[T]) Get() T {
	var ret T
	if !n.IsEmpty() {
		ret = n.value
	}
	return ret
}

func (n *node[T]) Rest() types.List[T] {
	return n.next
}

func (n *node[T]) Equals(o types.List[T]) bool {
	if n.Count() == o.Count() {
		if n == o {
			return true
		}
		if n.Get() != o.Get() {
			return false
		}
		nIterator := types.NewListIterator(n.next)
		oIterator := types.NewListIterator(o.Rest())
		for {
			if !nIterator.HasNext() || !oIterator.HasNext() {
				break
			}
			if nIterator.Next() != oIterator.Next() {
				return false
			}
		}
		return true
	}
	return false
}

func (n *node[T]) ContainsValue(v T) bool {
	if n == nil || n.IsEmpty() {
		return false
	}

	if v == n.Get() {
		return true
	}

	it := types.NewListIterator(n.next)
	for it.HasNext() {
		if v == it.Next() {
			return true
		}
	}

	return false
}
