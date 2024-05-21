package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, job)
        time.Sleep(time.Second) // Simulate time-consuming task
        fmt.Printf("Worker %d finished job %d\n", id, job)
        results <- job * 2 // Send result back to results channel
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start three worker goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs to the jobs channel
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs) // Close the jobs channel to indicate no more jobs

    // Collect results from the results channel
    for a := 1; a <= numJobs; a++ {
        select {
        case result := <-results:
            fmt.Printf("Received result: %d\n", result)
        case <-time.After(2 * time.Second):
            fmt.Println("Timeout occurred")
        }
    }
}
