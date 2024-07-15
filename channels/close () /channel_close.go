// Go program to illustrate how 
// to close a channel using for 
// range loop and close function 
package main 

import "fmt"

// Function 
func myfun(mychnl chan string) { 

	for v := 0; v < 4; v++ { 
		mychnl <- "GeeksforGeeks"
	} 
	close(mychnl) 
} 

// Main function 
func main() { 

	// Creating a channel 
	c := make(chan string) 

	// calling Goroutine 
	go myfun(c) 

	// When the value of ok is 
	// set to true means the 
	// channel is open and it 
	// can send or receive data 
	// When the value of ok is set to 
	// false means the channel is closed 
	for res := range c {
			fmt.Println("Channel Open ", res) 
		
		}

	
		fmt.Println("Channel Close ") 
		
	} 
 
// In Go, the close() function is used to close a channel when no more values will be sent on it.
// Closing a channel is an important operation in concurrent programming because it informs all 
// goroutines that are receiving from the channel that no more values will be sent, allowing them 
// to finish their work and exit gracefully.