package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchMatrix(t *testing.T) {
	testCases := []struct {
		desc   string
		matrix [][]int
		target int
		expect bool
	}{
		{
			desc:   "TC0: it should return true",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			expect: true,
		}, {
			desc:   "TC1: it should return false",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			expect: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := searchMatrix(tC.matrix, tC.target)
			require.Equal(t, tC.expect, actual)
		})
	}
}
