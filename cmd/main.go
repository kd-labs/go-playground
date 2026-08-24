package main

func searchRange(nums []int, target int) []int {
	idx := binarySearch(nums, target, 0, len(nums)-1)
	if idx == -1 {
		return []int{-1, -1}
	}

	first, last := idx, idx
	res := make([]int, 2)

	for first != -1 {
		res[0] = first
		first = binarySearch(nums, target, 0, first-1)
	}

	for last != -1 {
		res[1] = last
		last = binarySearch(nums, target, last+1, len(nums)-1)
	}

	return res
}

func binarySearch(nums []int, target, lo, hi int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
