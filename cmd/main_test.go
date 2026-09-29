package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchInsert(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect int
	}{
		{
			desc:   "TC0: it should return 4",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 5,
			expect: 4,
		}, {
			desc:   "TC1: it should return 6",
			nums:   []int{-1, 0, 2, 4, 6, 8},
			target: 10,
			expect: 6,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchInsert(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}
