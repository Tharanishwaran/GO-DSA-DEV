package main

import (
   
	"fmt"
)

func main() {

fmt.Println("Hello, World!")

numtotal := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) 

fmt.Println(numtotal)

}

func sum(numbers ...int) int {

  total := 0 
  for _, num := range numbers {
  
	total += num
  }
  return total

}