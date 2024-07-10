package main

import "fmt"



func main()  {
	
b :=  factorial(5)

fmt.Println(b)


}

func factorial(a int) int {


	if a==0 || a==1{

		return 1;
	}else if a < 0{

      return 0;

	}
    
	return a * factorial(a-1) ;


}