package main

import (
    "fmt"
    "time"
)

// Function to send messages
func sender(ch chan int, msg int) {
    fmt.Println("Sending", msg)
    ch <- msg
    fmt.Println("Sent", msg)
}

func main() {
    // Buffered channel with a capacity of 5
    buffered := make(chan int, 5)

    // Send 2 messages to the buffered channel
    go sender(buffered, 1)
    go sender(buffered, 2)

    // Allow some time for senders to send messages
    time.Sleep(time.Second)

    // Receiver can receive messages from the buffer
    fmt.Println("Receiving buffered")
    fmt.Println("Received buffered", <-buffered)
    fmt.Println("Received buffered", <-buffered)
    fmt.Println("Done receiving buffered")
}
