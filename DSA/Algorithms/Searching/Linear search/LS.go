package main

import "fmt"

// LinearSearch function performs linear search on an array
func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // return the index of the found element
		}
	}
	return -1 // return -1 if the element is not found
}

func main() {
	// Test array
	arr := []int{3, 5, 7, 9, 11, 13, 15}

	// Element to search
	target := 9

	// Perform Linear Search
	index := LinearSearch(arr, target)

	if index != -1 {
		fmt.Printf("Element found at index: %d\n", index)
	} else {
		fmt.Println("Element not found in array")
	}
}
