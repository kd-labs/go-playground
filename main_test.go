package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapify(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "basic unsorted array",
			input:  []int{4, 10, 3, 5, 1},
			output: []int{1, 4, 3, 5, 10},
		},
		{
			name:   "descended sorting (worst case)",
			input:  []int{9, 7, 5, 3, 1},
			output: []int{1, 3, 5, 9, 7},
		},
		{
			name:   "single element",
			input:  []int{42},
			output: []int{42},
		},
		{
			name:   "duplicates",
			input:  []int{3, 3, 3, 1, 1},
			output: []int{1, 1, 3, 3, 3},
		},
		{
			name:   "empty arr",
			input:  []int{},
			output: []int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			BuildMinHeap(tC.input)
			require.Equal(t, tC.output, tC.input)
		})
	}
}
