package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleNonDuplicate(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "TC0: it should return 2",
			nums:   []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expect: 2,
		}, {
			desc:   "TC1: it should return 10",
			nums:   []int{3, 3, 7, 7, 10, 11, 11},
			expect: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := singleNonDuplicate(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}
