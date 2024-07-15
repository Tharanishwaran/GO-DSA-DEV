package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a new channel of type int
    ch := make(chan int)

    // Start a new goroutine that sends data to the channel
    go func() {
        fmt.Println("Sending data to the channel...")
        ch <- 42  // Send the value 42 to the channel
        fmt.Println("Data sent to the channel")
    }()

    // Simulate some work with a delay
    time.Sleep(1 * time.Second)

    // Receive data from the channel
    fmt.Println("Receiving data from the channel...")
    value := <-ch  // Receive the value from the channel
    fmt.Println("Data received from the channel:", value)
}
