package main

import (
	"bufio"
	"fmt"
	"os"
  )

func main(){

var name string

// method 1 fmt.Scanln

fmt.Print("enter your name: ")

fmt.Scanln(&name) // Address-of operator (&) is used to pass the memory address of the variable

fmt.Println("Hello,", name)

// method 2  scanf

var age int

  fmt.Println("Enter your name and age (separated by space):")
  fmt.Scanf("%s %d", &name, &age) // Format specifiers for string and integer
  fmt.Println("Hello,", name, "you are", age, "years old.")




// method 3



  reader := bufio.NewReader(os.Stdin) // Create reader object from standard input
  fmt.Println("Enter your name:")
  name1, _ := reader.ReadString('\n') // Read a line of input until newline character
  fmt.Println("Hello,", name1[:len(name1)-1]) // Remove trailing newline character

}





















