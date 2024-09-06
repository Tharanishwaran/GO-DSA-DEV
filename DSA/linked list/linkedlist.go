package main

import (
	"fmt"
)

// Node represents each element in the linked list
type Node struct {
	data int   // Change `int` to any type you want to store
	next *Node // Pointer to the next node in the list
}

// LinkedList manages the list operations
type LinkedList struct {
	head *Node // Pointer to the first node in the list
}

// Insert inserts a new node with the given data at the beginning of the list
func (ll *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	// If the list is empty, new node becomes the head
	if ll.head == nil {
		ll.head = newNode
	} else {
		// Otherwise, insert new node before the current head
		newNode.next = ll.head
		ll.head = newNode
	}
}

// Display prints all elements in the linked list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := LinkedList{}

	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)

	fmt.Println("Linked List:")
	ll.Display()
}
