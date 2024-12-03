package main

import (
    "fmt"
)

// Function to send messages
func sender(ch chan int, msg int) {
    fmt.Println("Sending", msg)
    ch <- msg
    fmt.Println("Sent", msg)
}

func main() {
    // Buffered channel with a capacity of 2
    buffered := make(chan int, 2)

    // Sending three messages but the buffer can only hold two
    go sender(buffered, 1)
    go sender(buffered, 2)
    go sender(buffered, 3)

    // Trying to receive only two messages
    fmt.Println("Receiving buffered")
    fmt.Println("Received buffered", <-buffered)
    fmt.Println("Received buffered", <-buffered)

    // The program will hang here because the buffer is full and there's no receiver
    // to receive the third message, causing a deadlock.
    fmt.Println("Done receiving buffered")
}
