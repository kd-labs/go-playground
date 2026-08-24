package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect []int
	}{
		{
			desc:   "it shold return [3,4]",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			expect: []int{3, 4},
		}, {
			desc:   "it shold return [-1,-1]",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			expect: []int{-1, -1},
		}, {
			desc:   "it shold return [-1,-1]",
			nums:   []int{},
			target: 0,
			expect: []int{-1, -1},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchRange(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}
