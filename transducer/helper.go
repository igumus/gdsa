package transducer

import (
	"fmt"
	"strings"
)

// predicates
var (
	IsEven = func(item int) bool {
		return item%2 == 0
	}

	IsOdd = func(item int) bool {
		return !IsEven(item)
	}

	LessThan = func(val int) func(int) bool {
		return func(i int) bool {
			return i < val
		}
	}

	GreaterThan = func(val int) func(int) bool {
		return func(i int) bool {
			return i > val
		}
	}
)

// maps/converters
var (
	IntIncrementer = func(item int) int {
		return item + 1
	}

	IntDecrementer = func(item int) int {
		return item - 1
	}

	IntStringfy = func(item int) string {
		return fmt.Sprintf("%d", item)
	}
)

// reducing functions
var (
	StringAppend = func(seperator string) ReducerFunction[string] {
		builder := strings.Builder{}
		return func(ret any, reduced *bool, items ...string) any {
			acc := ""
			if len(items) > 0 && !*reduced {
				builder.WriteString(items[0])
				builder.WriteString(seperator)
				acc = builder.String()
			} else {
				acc = builder.String()
				builder.Reset()
			}
			return acc
		}
	}
)
