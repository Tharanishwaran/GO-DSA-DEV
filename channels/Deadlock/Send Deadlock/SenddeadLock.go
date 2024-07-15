// Example Without Receiving Data

package main

import "fmt"

func main() {
    ch := make(chan int)

    // Sending data to the channel
    ch <- 42  // This will cause a deadlock because no goroutine is receiving the data

    fmt.Println("This line will never be reached")


	
}
