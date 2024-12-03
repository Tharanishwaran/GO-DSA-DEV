package main

import (
	"errors"
	"fmt"
)

func main() {


fmt.Println("Hello, World!")

result, err := divide(10, 2)

fmt.Println(result, err)

}


func divide(x ,y float64) (float64, error){

	if y == 0 {
         return 0, errors.New("cannot divide by zero")
	}

     return x/y,nil

}