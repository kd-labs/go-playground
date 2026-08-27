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
			desc: "TC1: it should return true",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target: 10,
			expect: true,
		}, {
			desc: "TC1: it should return false",
			matrix: [][]int{
				{1, 2, 4, 8},
				{10, 11, 12, 13},
				{14, 20, 30, 40},
			},
			target: 15,
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
