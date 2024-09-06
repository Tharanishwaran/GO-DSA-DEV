package main

import "fmt"

func BinarySearch(arr []int, target int) int {

	low, high := 0, len(arr)-1

	for low <= high {

		mid := low + (high-low)/2

		//check if target is present at mid
		if arr[mid] == target {
			return mid

		}
		//if target is greater, ignore the left half
		if arr[mid] < target {

			low = mid + 1

		} else {

			//if target is smaller,ignore the right half
			high = mid - 1

		}

	}

	return -1

}

func main() {

	arr := []int{10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70}
	target := 25

	result := BinarySearch(arr, target)

	if result != -1 {

		fmt.Println("Element found at index %d\n", result)

	} else {

		fmt.Println("Element not found")
	}

}
