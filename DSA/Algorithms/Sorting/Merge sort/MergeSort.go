package main

import (
	"fmt"
)

// Merge function to merge two sorted subarrays
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare elements of both subarrays and merge them in sorted order
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append the remaining elements of the left subarray
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Append the remaining elements of the right subarray
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// MergeSort function to recursively sort the array
func mergeSort(arr []int) []int {
	// Base case: if the array has 1 or no elements, it's already sorted
	if len(arr) < 2 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// Merge the sorted halves
	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
