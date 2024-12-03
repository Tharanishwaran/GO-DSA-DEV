package main

import "fmt"


func main() {
    ch := make(chan string)
    go func() {
        ch <- "Hello from Goroutine"
    }()
   // Receive data from the channel

	ch1 := make(chan string)
	go func() {
		ch1 <- "Hello from another goroutine"
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch)
}
