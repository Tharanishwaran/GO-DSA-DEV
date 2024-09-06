package main

import (
    "fmt"
    "time"
)

// Function to send a message
func sender(ch chan int, msg int) {
    fmt.Println("Sending", msg)
    ch <- msg
    fmt.Println("Sent", msg)
}

func main() {
    // Unbuffered channel
    unbuffered := make(chan int)

    // Start a goroutine to send a message
    go sender(unbuffered, 1)

    // Sleep for a second to show the sender is waiting
    time.Sleep(time.Second)

    // Main goroutine will receive the message
    fmt.Println("Receiving")
    msg := <-unbuffered
    fmt.Println("Received", msg)
}
