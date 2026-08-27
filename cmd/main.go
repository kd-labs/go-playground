package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	lo, hi := 0, rows*cols-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		r := mid / cols
		c := mid % cols

		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}
