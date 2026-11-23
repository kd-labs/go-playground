package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLowerBound(t *testing.T) {
	testCases := []struct {
		desc string
		nums []int
		x    int
		want int
	}{
		{
			desc: "TC1: it should return index 1",
			nums: []int{1, 2, 2, 3},
			x:    2,
			want: 3,
		}, {
			desc: "TC2: it should return index 3",
			nums: []int{3, 5, 8, 15, 19},
			x:    9,
			want: 3,
		}, {
			desc: "TC3: it should return index 3",
			nums: []int{3, 5, 8, 15, 15},
			x:    9,
			want: 3,
		}, {
			desc: "TC4: it should return index 5",
			nums: []int{3, 5, 8, 8, 19},
			x:    8,
			want: 4,
		}, {
			desc: "TC5: it should return index 5",
			nums: []int{3, 5, 8, 8, 19},
			x:    19,
			want: 5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := upperBound(tC.nums, tC.x)
			require.Equal(t, tC.want, got)
		})
	}
}
