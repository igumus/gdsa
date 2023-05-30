package transducer

import (
	"testing"

	"github.com/igumus/gdsa/types"
	"github.com/stretchr/testify/assert"
)

func toArray[T any](it types.Iterator[T]) []T {
	ret := make([]T, 0)
	for it.HasNext() {
		ret = append(ret, it.Next())
	}
	return ret
}

func TestFunctionMap(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(10))
	tf1 := Map(IntStringfy)
	xf := tf1(StringAppend(" "))
	output := ""
	output = Reduce(xf, output, coll)
	assert.Equal(t, "0 1 2 3 4 5 6 7 8 9 ", output)

	coll = types.NewFiniteRange(types.WithEnd(10))
	tf2 := Map(IntIncrementer)
	xf2 := tf2(Append[int])
	output2 := make([]int, 0)
	output2 = Reduce(xf2, output2, coll)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, output2)
}

func TestFunctionFilter(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(10))
	tf := Filter(IsOdd)
	xf := tf(Append[int])
	output := make([]int, 0)
	output = Reduce(xf, output, coll)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, output)
}

func TestFunctionCombine(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(10))
	tf := Combine(Filter(IsOdd), Map(IntStringfy))
	xf := tf(Append[string])
	output := make([]string, 0)
	output = Reduce(xf, output, coll)
	assert.Equal(t, []string{"1", "3", "5", "7", "9"}, output)
}

func TestFunctionTake(t *testing.T) {
	coll := types.NewInfiniteRange()
	xf := Take[int](5)
	output := make([]int, 0)
	tf := xf(Append[int])
	output = Reduce(tf, output, coll)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, output)
}

func TestFunctionTakeWhile(t *testing.T) {
	coll := types.NewInfiniteRange()
	sampleSize := 5
	xf := TakeWhile(LessThan(sampleSize))
	output := make([]int, 0)
	tf := xf(Append[int])
	output = Reduce(tf, output, coll)
	assert.Equal(t, []int{0, 1, 2, 3, 4}, output)
}

func TestFunctionSkip(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(10))
	tf := Skip[int](5)
	xf := tf(Append[int])
	output := make([]int, 0)
	output = Reduce(xf, output, coll)
	assert.Equal(t, []int{5, 6, 7, 8, 9}, output)
}

func TestFunctionSkipWhile(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(20))
	tf := SkipWhile(LessThan(10))
	xf := tf(Append[int])
	output := make([]int, 0)
	output = Reduce(xf, output, coll)
	assert.Equal(t, []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}, output)
}

func TestFunctionEveryNth(t *testing.T) {
	coll := types.NewFiniteRange(types.WithEnd(10))
	tf := EveryNth[int](2)
	xf := tf(Append[int])
	output := make([]int, 0)
	output = Reduce(xf, output, coll)
	expected := []int{0, 2, 4, 6, 8}
	assert.Equal(t, expected, output)
}

func TestFunctionPartitionAll(t *testing.T) {
	tf := PartitionAll[int](2)
	xf := tf(Append[types.Iterator[int]])
	output := make([]types.Iterator[int], 0)
	coll := types.NewFiniteRange(types.WithEnd(10))
	output = Reduce(xf, output, coll)
	assert.Equal(t, 5, len(output))

	oddColl := types.NewFiniteRange(types.WithEnd(11))
	output = make([]types.Iterator[int], 0)
	output = Reduce(xf, output, oddColl)
	assert.Equal(t, 6, len(output))
}

func TestFunctionPartitionBy(t *testing.T) {
	fn := func(item int) int {
		return item
	}
	tf := PartitionBy(fn)
	xf := tf(Append[types.Iterator[int]])
	output := make([]types.Iterator[int], 0)
	coll := types.NewSliceIterator([]int{1, 1, 1, 2, 2, 3, 4, 5, 5})
	expected := [][]int{
		{1, 1, 1},
		{2, 2},
		{3},
		{4},
		{5, 5},
	}
	output = Reduce(xf, output, coll)
	for i := 0; i < len(output); i++ {
		it := output[i]
		arr := toArray(it)
		assert.Equal(t, expected[i], arr)
	}
	assert.Equal(t, 5, len(output))
}
