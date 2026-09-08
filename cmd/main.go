package main

func search(nums []int, target int) bool {
	pivot := minIndex(nums)

	var ans int
	if nums[pivot] == target {
		return true
	} else if nums[pivot] < target && target <= nums[len(nums)-1] {
		ans = binarySearch(nums, pivot+1, len(nums)-1, target)
	} else {
		ans = binarySearch(nums, 0, pivot-1, target)
	}
	return ans > -1
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

	for lo < hi && nums[lo] == nums[lo+1] {
		lo++
	}

	for lo < hi && nums[hi] == nums[hi-1] {
		hi--
	}

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
