package main

import (
    "fmt"
    "sync"
	"time"
)

var (
    count int
    mu    sync.Mutex
)

func increment() {
    mu.Lock()
    count++
    fmt.Println("Incremented count to:", count)
    mu.Unlock()
}

func main() {
    numIncrements := 5
    for i := 0; i < numIncrements; i++ {
        go increment()
    }

    // Wait for a short time to allow goroutines to finish (not a good practice for real use cases)
    time.Sleep(time.Second)
    fmt.Println("Final count:", count)
}
