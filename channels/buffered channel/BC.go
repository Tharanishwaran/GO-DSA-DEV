package main

import "fmt"

func main() {
    // Creating a buffered channel with a capacity of 3
    ch := make(chan int, 3)

    // Sending values to the buffered channel
    ch <- 1
    ch <- 2
    ch <- 3
	// ch <-4

    // Since the buffer is full, the next send operation will block
    // Uncommenting the following line will cause a deadlock if there is no corresponding receive operation
    // ch <- 4

    // Receiving values from the buffered channel
    fmt.Println(<-ch) // Outputs: 1
    fmt.Println(<-ch) // Outputs: 2
	fmt.Println(<-ch) // Outputs: 3
	// fmt.Println(<-ch) // 

    // Now the buffer is empty, so the next receive operation will block
    // Uncommenting the following line will cause a deadlock if there is no corresponding send operation
    // fmt.Println(<-ch)
}
