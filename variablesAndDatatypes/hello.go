package main

import "fmt"

import "unicode/utf8"

func main() {

var intnum int = 32767
  
intnum += 1
  
fmt.Println(intnum)
  
fmt.Println("Hello, world!")

var h float32 = 4.0

var h1 float32 =  7.9

var result int64  = int64(h) + int64(h1)

fmt.Println(result)

var mystring string = `hello   
world`  // ` for multi - line strings

var mystring1 string = "hello universe" //" for singline strings"

fmt.Println(mystring)

fmt.Println(mystring1)

fmt.Println(len("A")) //The len() method returns the number of bytes in a string

fmt.Println(utf8.RuneCountInString(mystring1))

var myrune rune = 'a'

fmt.Println(myrune)


var mybol bool= true

fmt.Println(mybol)

var intnum3 = 12 //we can intiialize variables without declare its data type

fmt.Println(intnum3)

myvar := 12 //we can also intialize a variable without var keyword

var1 , var2 := 10,20

fmt.Println(myvar)

fmt.Println(var1,var2)

const cannotchange = "i am never change"



hello := "hello everyone"

println(hello)

println(add(1,2))


}

func add(a int,b int) int { //int in the outside of the () means return type

 return a+b

}


