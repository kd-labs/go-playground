package main

func findKthPositive(arr []int, k int) int {
	arrSize := len(arr)
	nums := make([]int, arrSize)

	for i := 0; i < arrSize; i++ {
		nums[i] = i + 1
	}

	left, right := 0, arrSize-1

	for left <= right {
		mid := left + (right-left)/2

		misses := arr[mid] - nums[mid]
		if misses < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left + k
}
