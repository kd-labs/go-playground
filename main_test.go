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

func TestExtractMin(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		min    int
		output []int
		err    error
	}{
		{
			name:   "general case",
			input:  []int{1, 4, 3, 5, 10},
			output: []int{3, 4, 10, 5},
			min:    1,
			err:    nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			min, arr, err := ExtractMin(tC.input)
			require.Equal(t, tC.min, min)
			require.Equal(t, tC.output, arr)
			require.Equal(t, tC.err, err)
		})
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		value  int
		output []int
	}{
		{
			name:   "happy path",
			input:  []int{1, 4, 3, 5, 10},
			value:  2,
			output: []int{1, 4, 2, 5, 10, 3},
		},
	}

	for _, tC := range tests {
		t.Run(tC.name, func(t *testing.T) {
			actual := Insert(tC.input, tC.value)
			require.Equal(t, tC.output, actual)
		})
	}
}
