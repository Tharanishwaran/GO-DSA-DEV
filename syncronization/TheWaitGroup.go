package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done() // Decrement the counter when the goroutine completes
    fmt.Printf("Worker %d starting\n", id)
    // Simulate some work
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    numWorkers := 3
    wg.Add(numWorkers) // Set the number of goroutines to wait for

    for i := 1; i <= numWorkers; i++ {
        go worker(i)
    }

    wg.Wait() // Block until all goroutines have finished
    fmt.Println("All workers done")
}
