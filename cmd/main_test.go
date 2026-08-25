package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "TC1: it should return 1",
			nums:   []int{3, 4, 5, 1, 2},
			expect: 1,
		}, {
			desc:   "TC2: it should return 0",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			expect: 0,
		}, {
			desc:   "TC3: it should return 11",
			nums:   []int{11, 13, 15, 17},
			expect: 11,
		}, {
			desc:   "TC4: it should return 1",
			nums:   []int{2, 1},
			expect: 1,
		}, {
			desc:   "it should return 1",
			nums:   []int{3, 1, 2},
			expect: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findMin(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}
