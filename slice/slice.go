package main

import "fmt"

func main() {
    // Declaring a slice
    var numbers []int

    // Creating a slice using the make function
    numbers = make([]int, 5) // A slice with length 5 and capacity 5

    // Creating a slice from an array
    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4] // A slice from index 1 to 3 (4 is exclusive)

    // Appending elements to a slice
    numbers = append(numbers, 6, 7, 8)

    // Accessing elements in a slice
    fmt.Println(slice[0]) // Output: 2

    // Getting length and capacity of a slice
    fmt.Println(len(slice)) // Output: 3
    fmt.Println(cap(slice)) // Output: 4

    // Iterating over a slice using a for loop
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }

    // Iterating over a slice using a range loop
    for index, value := range slice {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Slicing a slice
    subSlice := slice[1:3] // A new slice from index 1 to 2
    fmt.Println(subSlice)  // Output: [3 4]
}



// 1. Variables and Constants
// 2. Data Types
// 3. Operators
// 4. Control Structures (if, else, switch, for, break, continue)
// 5. Arrays and Slices
// 6. Maps
// 7. Structs
// 8. Pointers
// 9. Functions
// 10. Methods
// 11. Interfaces
// 12. Goroutines
// 13. Channels
// 14. Packages and Imports
// 15. Error Handling
// 16. Defer
// 17. Panic and Recover
// 18. Reflection
// 19. Testing
// 20. Concurrency Patterns