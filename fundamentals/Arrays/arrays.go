package main

import "fmt"

func main() {
    // Declaring an array of integers with a size of 5
    var numbers []int

    // Initializing the array with values
    numbers = []int{1, 2, 3, 4, 5}

    // Accessing and printing the first element of the array
    // fmt.Println(numbers[0]) // Output: 1

    // Getting the length of the array
    fmt.Println(len(numbers) -1 ) // Output: 5



    // Iterating over the array and printing each element
    for i := 0; i <= len(numbers) -1; i++ {
        fmt.Println(numbers[i])
    }

    // Declaring and initializing a 2D array
    // var matrix [3][3]int

    // matrix := [3][3]int{
    //     {1, 2, 3},
    //     {4, 5, 6},
    //     {7, 8, 9},
    // }

	// fmt.Println(matrix)
}
