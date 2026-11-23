package main

func upperBound(nums []int, x int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > x {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
