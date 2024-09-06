package main

import "fmt"
import "runtime"

func main()  {

// if-else
x := 10

if x*10 == 100{

fmt.Println("x is 100")

} else { //else keyword always start at end of the } brakets

	fmt.Println("x is not 100")
}



// switch statement

day := "Monday"

switch day {
  case "Saturday", "Sunday":
    fmt.Println("Weekend!")
  case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
  default:
    fmt.Println("Invalid day")
}

fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os) 
  }




}