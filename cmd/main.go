package main

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2

		r := mid / cols
		c := mid % cols

		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
