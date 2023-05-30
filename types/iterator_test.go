package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringIterator(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty iterator has next",
			input:    "",
			expected: false,
		},
		{
			name:     "one item slice iterator has next",
			input:    "1",
			expected: true,
		},
		{
			name:     "two item slice iterator has next",
			input:    "12",
			expected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			iterator := NewStringIterator(tc.input)
			assert.Equal(t, tc.expected, iterator.HasNext())
		})
	}

	source := "hello"
	iterator := NewStringIterator(source)
	for _, item := range source {
		assert.Equal(t, true, iterator.HasNext())
		assert.Equal(t, iterator.Next(), item)
	}
}

func TestSliceIterator(t *testing.T) {
	testcases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "nil iterator has next",
			input:    nil,
			expected: false,
		},
		{
			name:     "empty iterator has next",
			input:    []int{},
			expected: false,
		},
		{
			name:     "one item slice iterator has next",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "two item slice iterator has next",
			input:    []int{1, 2},
			expected: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			iterator := NewSliceIterator(tc.input)
			assert.Equal(t, tc.expected, iterator.HasNext())
		})
	}

	source := []int{1, 2, 3}
	iterator := NewSliceIterator(source)
	for _, item := range source {
		assert.Equal(t, true, iterator.HasNext())
		assert.Equal(t, iterator.Next(), item)
	}
	assert.Equal(t, false, iterator.HasNext())
}

func TestRangeIterator(t *testing.T) {
	end := 10
	it := NewFiniteRange(WithEnd(end))
	i := 0
	for it.HasNext() {
		assert.Equal(t, i, it.Next())
		i++
	}
}
