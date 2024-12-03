package main

import (
    "fmt"
    "sync"
    "time"
)

var (
    cond  = sync.NewCond(&sync.Mutex{}) // Initialize a condition variable with a mutex
    ready bool                          // Shared variable to indicate readiness
)

func waitForSignal() {
    cond.L.Lock()                // Lock the mutex before checking the condition
    for !ready {                 // Loop until the condition becomes true
        cond.Wait()              // Wait releases the lock and waits for a signal
    }
    fmt.Println("Received signal") // Condition met, print message
    cond.L.Unlock()              // Unlock the mutex
}

func main() {
    go waitForSignal()            // Start a new goroutine to wait for the signal

    time.Sleep(time.Second)       // Simulate some work in the main goroutine

    cond.L.Lock()                 // Lock the mutex before changing the condition
    ready = true                  // Set the condition to true
    cond.Signal()                 // Signal the waiting goroutine that the condition is met
    cond.L.Unlock()               // Unlock the mutex

    time.Sleep(time.Second)       // Wait to ensure the signal is processed
}
