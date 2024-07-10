package main

import "fmt"

func main()  {
	
a := 0;
b := 0;

fmt.Println("enter value for A")

fmt.Scan(&a)

fmt.Println("enter value for B")

fmt.Scan(&b)

if  a > b{
	

	fmt.Println("a is greater")
}else{

   fmt.Println("b is greater")

}

	
}