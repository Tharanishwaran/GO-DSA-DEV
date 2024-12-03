package main

import (
    "fmt"
	"time"
)


func sayHello() {
    fmt.Println("Hello, World!")
}

func main() {
    go sayHello() // Run concurrently
    time.Sleep(time.Second) // Allow goroutine to complete
	fmt.Println("Done")
}
