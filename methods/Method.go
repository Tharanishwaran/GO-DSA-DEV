// In Go, methods are functions that are associated with a specific type. 
// They are defined with a special receiver argument that binds the method to the type.
//  Methods allow you to define behavior associated with types, 
//  providing a way to work with data encapsulation and type-specific functionality.



package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Method with a value receiver
func (p Person) greet() string {
    return "Hello, " + p.Name
}

func main() {
    // Create an instance of the Person struct
    person := Person{Name: "Alice", Age: 30}
    
    // Call the method on the instance
    message := person.greet()
    fmt.Println(message) // Output: Hello, Alice
}
