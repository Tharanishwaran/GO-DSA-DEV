package main

import "fmt"

func main() {
    ch := make(chan int)

    // Receiving data from the channel
    value := <-ch  // This will cause a deadlock because no goroutine is sending data

    fmt.Println("Received:", value)
}


