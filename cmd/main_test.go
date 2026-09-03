package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindKRotation(t *testing.T) {
	testCases := []struct {
		desc   string
		nums   []int
		expect int
	}{
		{
			desc:   "TC0: it should return 1",
			nums:   []int{5, 1, 2, 3, 4},
			expect: 1,
		}, {
			desc:   "TC0: it should return 0",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := findKRotation(tC.nums)
			require.Equal(t, tC.expect, actual)
		})
	}
}
