package list

import (
	"github.com/igumus/gdsa/types"
)

type list[T comparable] struct {
	root types.List[T]
}

func NewList[T comparable](mutable bool) types.List[T] {
	var ret types.List[T] = emptyNode[T]()
	if mutable {
		ret = &list[T]{
			root: nil,
		}
	}

	return ret
}

func NewListFromArray[T comparable](mutable bool, items []T) types.List[T] {
	var ret types.List[T] = newListFromArray(items)
	if mutable {
		ret = &list[T]{
			root: newListFromArray(items),
		}
	}

	return ret
}

func (l *list[T]) IsEmpty() bool {
	return l == nil || l.Count() == 0
}

func (l *list[T]) Count() int {
	if l == nil || l.root == nil {
		return 0
	}
	return l.root.Count()
}

func (l *list[T]) Get() T {
	return l.root.Get()
}

func (l *list[T]) Rest() types.List[T] {
	if l == nil || l.root == nil {
		return nil
	}
	return l.root.Rest()
}

func (l *list[T]) Add(value T) types.List[T] {
	if l.root == nil {
		l.root = newNode(value)
	} else {
		nroot, _ := l.root.Add(value).(*node[T])
		l.root = nroot
	}
	return l
}

func (l *list[T]) ContainsValue(v T) bool {
	if l == nil || l.root == nil {
		return false
	}
	return l.root.ContainsValue(v)
}

func (l *list[T]) Equals(o types.List[T]) bool {
	if otherList, ok := o.(*list[T]); ok {
		// o is type of list
		if l == otherList {
			return true
		}
		return l.root.Equals(otherList.root)
	}
	// other is type of node
	return l.root.Equals(o)
}
