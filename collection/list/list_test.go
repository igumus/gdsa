package list

import (
	"testing"

	"github.com/igumus/gdsa/transducer"
	"github.com/igumus/gdsa/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListImmutable(t *testing.T) {
	list := NewList[int](false)
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())
	other := list.Add(1)

	require.NotEqual(t, list, other)
	require.False(t, other.IsEmpty())
	require.Equal(t, 1, other.Count())
}

func TestListMutable(t *testing.T) {
	list := NewList[int](true)
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())
	other := list.Add(1)

	require.Equal(t, list, other)
	require.False(t, other.IsEmpty())
	require.Equal(t, 1, other.Count())
}

func TestListCreate(t *testing.T) {
	var list *list[int] = nil
	var iterator types.Iterator[int]
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())

	other := NewList[int](false)
	iterator = types.NewListIterator(other)
	require.True(t, other.IsEmpty())
	require.Equal(t, 0, other.Count())
	require.False(t, iterator.HasNext())

	nilArray := NewListFromArray[int](false, nil)
	require.True(t, nilArray.IsEmpty())
	require.Equal(t, 0, nilArray.Count())
	iterator = types.NewListIterator(nilArray)
	require.False(t, iterator.HasNext())

	emptyArray := NewListFromArray(false, []int{})
	require.True(t, emptyArray.IsEmpty())
	require.Equal(t, 0, emptyArray.Count())
	iterator = types.NewListIterator(emptyArray)
	require.False(t, iterator.HasNext())

	another := NewListFromArray(false, []int{1, 2, 3})
	require.False(t, another.IsEmpty())
	require.Equal(t, 3, another.Count())
	require.Equal(t, 1, another.Get())
	iterator = types.NewListIterator(another)
	require.True(t, iterator.HasNext())
}

func TestListAddElement(t *testing.T) {
	list := NewList[int](true)
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())
	list.Add(1)
	require.False(t, list.IsEmpty())
	require.Equal(t, 1, list.Count())
	require.Equal(t, 1, list.Get())
	list.Add(2)
	require.False(t, list.IsEmpty())
	require.Equal(t, 2, list.Count())
	require.Equal(t, 2, list.Get())
}

func TestListRest(t *testing.T) {
	list := NewList[int](true)
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())
	require.Nil(t, list.Rest())

	list.Add(1)
	rest := list.Rest()
	require.Nil(t, rest)

	other := list.Add(2)
	require.Equal(t, other, list)
	rest = list.Rest()
	require.NotNil(t, rest)
	require.False(t, rest.IsEmpty())
	require.Equal(t, 1, rest.Count())
	require.Equal(t, 1, rest.Get())
	require.Nil(t, rest.Rest())
}

func TestListContainsValue(t *testing.T) {
	list := NewList[int](true)
	require.True(t, list.IsEmpty())
	require.Equal(t, 0, list.Count())
	require.Nil(t, list.Rest())

	// check value on empty list
	require.False(t, list.ContainsValue(1))

	// check existence after adding element
	list.Add(1)
	require.True(t, list.ContainsValue(1))

	arr := []int{1, 2, 3, 4}
	list = NewListFromArray(true, arr)
	require.False(t, list.IsEmpty())
	require.Equal(t, 4, list.Count())
	require.NotNil(t, list.Rest())
	for _, i := range arr {
		require.True(t, list.ContainsValue(i))
	}
	require.False(t, list.ContainsValue(5))
}

func TestListMap(t *testing.T) {
	collIterator := types.NewListIterator(NewListFromArray(true, []int{1, 2, 3, 4}))
	tf1 := transducer.Map(transducer.IntIncrementer)
	xf := tf1(transducer.Append[int])
	output := NewList[int](true)
	output = transducer.Reduce(xf, output, collIterator)
	expected := NewListFromArray(true, []int{5, 4, 3, 2})
	assert.Equal(t, expected.Count(), output.Count())
	assert.True(t, output.Equals(expected))
}

func TestListFilter(t *testing.T) {
	collIterator := types.NewListIterator(NewListFromArray(true, []int{1, 2, 3, 4}))
	tf1 := transducer.Filter(transducer.IsEven)
	xf := tf1(transducer.Append[int])
	output := NewList[int](true)
	output = transducer.Reduce(xf, output, collIterator)
	expected := NewListFromArray(true, []int{4, 2})
	assert.Equal(t, expected.Count(), output.Count())
	assert.True(t, output.Equals(expected))
}
