package main

import "math"

func findMin(nums []int) int {
	res := math.MaxInt

	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)/2

		if nums[lo] <= nums[mid] {
			res = min(res, nums[lo])
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return min(res, nums[hi])
}
