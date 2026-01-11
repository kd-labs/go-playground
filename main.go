package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Running...")
}

// BuildMinHeap starts from last non-leaf node and heapifies down
func BuildMinHeap(arr []int) {
	indx := (len(arr) - 2) / 2

	for ; indx >= 0; indx-- {
		HeapifyDown(arr, indx, len(arr))
	}
}

// HeapifyDown takes value at indx i and puts in correct position in heap
func HeapifyDown(arr []int, i int, heapSize int) {
	current := i

	for {
		left := 2*current + 1
		right := 2*current + 2

		// assume smallest is current
		smallest := current

		if left < heapSize && arr[smallest] > arr[left] {
			smallest = left
		}

		if right < heapSize && arr[smallest] > arr[right] {
			smallest = right
		}

		// current node is in correct heap position
		if smallest == current {
			break
		}

		arr[current], arr[smallest] = arr[smallest], arr[current]

		// set current to smallest and heapify down
		current = smallest
	}
}

func ExtractMin(arr []int) (int, []int, error) {
	if len(arr) == 0 {
		return 0, nil, errors.New("empty array")
	}
	min := arr[0]
	heapSize := len(arr)

	arr[0] = arr[heapSize-1]

	// reduce heapsize
	heapSize = heapSize - 1

	HeapifyDown(arr, 0, heapSize)

	return min, arr[:heapSize], nil
}

func Insert(arr []int, value int) []int {
	newArr := make([]int, len(arr)+1)

	newArr = append(arr, value)
	HeapifyUp(newArr, len(newArr)-1)
	return newArr
}

func HeapifyUp(arr []int, i int) {
	for {
		parentIndx := (i - 1) / 2

		if arr[parentIndx] > arr[i] {
			arr[parentIndx], arr[i] = arr[i], arr[parentIndx]
			i = parentIndx
		} else {
			break
		}
	}
}
