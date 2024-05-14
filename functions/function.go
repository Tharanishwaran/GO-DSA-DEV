package main

import "fmt"

func main()  {
	
value := add(6,5)

fmt.Println(value)

greet("akash")

}

func add(x int, y int) int {
	return x + y
  }
  
  func greet(name string) {
	fmt.Println("Hello,", name)
  }

