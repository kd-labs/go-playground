package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		target int
		expect bool
	}{
		{
			desc:   "TC0: it should return true",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			expect: true,
		}, {
			desc:   "TC1: it should return false",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			expect: false,
		}, {
			desc:   "TC2: it should return true",
			nums:   []int{2, 2, 2, 3, 2, 2, 2},
			target: 3,
			expect: true,
		}, {
			desc:   "TC3: it should return true",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			target: 2,
			expect: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := search(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}
