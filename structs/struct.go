package main

import "fmt"

type point struct {
  x int
  y int
  c bool
}

type Person struct {
  Name string
  Age  int
  City string
}

func main() {
  // Create a Point using named field initialization
   a := point{ 10, 20, false}

   b := point{y:8} // x and c are now be an default values of int and bool

  // Print the Point with formatted output/
  fmt.Printf("Point a: x = %d, y = %d\n", a.x, a.y)

  fmt.Println(b)

 a1 := Person{Name:"akash",Age: 21,}

 fmt.Println(a1)




}

