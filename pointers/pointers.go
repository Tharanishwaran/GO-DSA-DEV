package main

import "fmt"


func main(){

var a = 10;

fmt.Println(a);

x := 7
y := &x

fmt.Println(x,y)
fmt.Println(*y) //it prints what value in the address

*y = 8 



fmt.Println(x,y)

	// Declare a variable and assign a value
	num := 42

	// Declare a pointer variable to hold the memory address of num
	var ptr *int = &num // "&" operator gets the address

	// Print the original value of num
	fmt.Println("Original value of num:", num)

	// Modify the value of num indirectly through the pointer
	*ptr = 100 // Dereference pointer and assign new value

	// Print the value of num again (it has changed)
	fmt.Println("Value of num after modification through pointer:", num)

	// Print the memory address stored in the pointer
	fmt.Println("Memory address stored in ptr:", ptr)

	// Another way to declare and initialize a pointer
	name := "Alice"
	newName := &name // Short declaration, pointer pointing to "Alice"

	// Modify the string value through the pointer (not recommended for strings)
	*newName = "Bob" // This might cause unexpected behavior with strings (explained below)

	// Print the modified name
	fmt.Println("Modified name:", name) // Name is still "Alice" (explained below)

    
	

	fmt.Println(a)

	changevalue(&a) 

	fmt.Println(a)




}


func changevalue(intnum *int){

	*intnum = 10

    
}

