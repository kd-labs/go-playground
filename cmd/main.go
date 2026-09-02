package main

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == nums[mid+1] {
			if (hi-mid+1)%2 != 0 {
				lo = mid + 2
			} else {
				hi = mid - 1
			}
		} else {
			if (hi-mid)%2 != 0 {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	}

	return nums[lo]
}
