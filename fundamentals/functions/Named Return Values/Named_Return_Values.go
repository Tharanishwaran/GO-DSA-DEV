package main

import "fmt"

func main() {

	fmt.Println("Hello, World!")

	fmt.Println(split(100))

}

func split(sum int) (x, y int) {

	x = sum * 4 / 10
	y = sum - x		 // naked return - returns x and y automatically

	return 
}