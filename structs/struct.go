package main

import "fmt"

func main() {
  // Create a Point using named field initialization
   a := point{ 10, 20}

   b := point{y:8} // x and c are now be an default values of int and bool

  // Print the Point with formatted output/
  fmt.Printf("Point a: x = %d, y = %d\n", a.x, a.y)

  fmt.Println(b)


 a1 := Person{Name:"akash",Age: 21,}

 fmt.Println(a1)

 p1 := point{y: 3}

 fmt.Println(p1)

fmt.Println(change2(p1))

 fmt.Println(p1)


Person1 := Person{"ak",6,"newyork"} 

Person1.Age =21

fmt.Println(Person1)

c1 := Circle{4.7,&point{4,8}}

fmt.Println(c1.point)




}

type point struct {
  x int
  y int
  // c bool
}

type Person struct {
  Name string
  Age  int
  City string
}

func change (p1 *point) {

  p1.y = 10


}

func change2 (p1 point) point{

  p1.y = 10

  return p1


}

type Circle struct{

radius float64

// centre *point // struct inside other struct //we can all use new variable for other struct

*point

}