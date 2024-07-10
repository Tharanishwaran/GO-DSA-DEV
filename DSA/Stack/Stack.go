package main

import "fmt"


type Stack struct{

 items [] int;
 

}




func (s* Stack) push(items int)  {
	

  s.items = append(s.items, items);

}


func (s* Stack ) pop() int  {
	
if s.IsEmpty(){
    
	fmt.Println("Stack is empty");
	return -1;
}

lastindex := len(s.items) -1;

lastitems := s.items[lastindex];

s.items = s.items[:lastindex];

return lastitems



}



func (s* Stack) IsEmpty() bool{

  return  len(s.items) == 0;

}


func (s* Stack) display()  {

   
   for _, item := range s.items{

      
       fmt.Print(item, "->")

   }


		
	}	







func main()  {
	
	stack := Stack{};

	stack.push(5);
	stack.push(10);
	stack.push(6);

	stack.display();



}