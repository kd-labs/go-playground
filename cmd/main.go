package main

func search(nums []int, target int) int {
	pivot := minIndex(nums)

	if nums[pivot] == target {
		return pivot
	} else if nums[pivot] < target && target <= nums[len(nums)-1] {
		return binarySearch(nums, pivot+1, len(nums)-1, target)
	} else {
		return binarySearch(nums, 0, pivot-1, target)
	}
}

func binarySearch(nums []int, lo, hi, target int) int {
	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func minIndex(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)/2

		if nums[mid] < nums[hi] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
