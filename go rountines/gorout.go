// Go program to illustrate
// the concept of Goroutine
package main

import (
	"fmt"
	"time"
)

// Go program to illustrate how 
// to create an anonymous Goroutine 
// Main function 
func main() { 

	fmt.Println("Welcome!! to Main function") 

	// Creating Anonymous Goroutine 
	go func() { 

		fmt.Println("Welcome!! to GeeksforGeeks") 
	}() 

	time.Sleep(1 * time.Second) 
	fmt.Println("GoodBye!! to Main function") 
} 





// func display(str string) { 
// 	for w := 0; w < 6; w++ { 
// 		fmt.Println(str) 

// 		time.Sleep(1 * time.Second) 

// 	} 
// } 

// func main() { 

// 	// Calling Goroutine 
//     go display("Welcome") 

// 	// Calling normal function 
// 	display("GeeksforGeeks") 
// } 

