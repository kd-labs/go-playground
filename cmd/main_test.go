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
		expect int
	}{
		{
			desc:   "TC0: it should return 4",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			expect: 4,
		}, {
			desc:   "TC1: it should return 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			expect: 1,
		}, {
			desc:   "TC2: it should return -1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			expect: -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := search(tC.nums, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}
