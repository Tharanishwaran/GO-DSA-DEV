package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
    perimeter() float64
}

type Triangle struct {
    base   float64
    height float64
    side1  float64
    side2  float64
}

func (t Triangle) area() float64 {
    return 0.5 * t.base * t.height
}

func (t Triangle) perimeter() float64 {
    return t.base + t.side1 + t.side2
}

type Rectangle struct {
    length float64
    width  float64
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

func (r Rectangle) perimeter() float64 {
    return 2*r.length + 2*r.width
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

func printDetails(s Shape) {
    fmt.Printf("Area: %f, Perimeter: %f\n", s.area(), s.perimeter())
}

func main() {
    triangle := Triangle{base: 3, height: 4, side1: 5, side2: 5}
    rectangle := Rectangle{length: 5, width: 4}
    circle := Circle{radius: 5}

    fmt.Println("Triangle:")
    printDetails(triangle)

    fmt.Println("Rectangle:")
    printDetails(rectangle)

    fmt.Println("Circle:")
    printDetails(circle)
}
