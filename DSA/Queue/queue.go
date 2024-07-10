package main

import (
    "fmt"
)

// Queue represents a queue data structure
type Queue struct {
    items    []int
    frontIdx int
}

// Enqueue adds an element to the back of the queue
func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}

// Dequeue removes an element from the front of the queue
// and returns it. If the queue is empty, it returns -1.
func (q *Queue) Dequeue() int {
    if q.IsEmpty() {
        fmt.Println("Queue is empty")
        return -1
    }
    frontItem := q.items[q.frontIdx]
    q.frontIdx++
    // Reset slice if all elements have been dequeued
    if q.frontIdx == len(q.items) {
        q.items = nil
        q.frontIdx = 0
    }
    return frontItem
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
    return q.frontIdx == len(q.items)
}

// Display prints all elements in the queue
func (q *Queue) Display() {
    for i := q.frontIdx; i < len(q.items); i++ {
        fmt.Print(q.items[i], " <- ")
    }
    fmt.Println()
}

func main() {
    queue := Queue{}

    queue.Enqueue(5)
    queue.Enqueue(10)
    queue.Enqueue(15)

    fmt.Println("Queue contents:")
    queue.Display()

    fmt.Println("Dequeued item:", queue.Dequeue())
    fmt.Println("Queue contents after dequeue:")
    queue.Display()
}
