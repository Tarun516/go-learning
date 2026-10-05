package main

import "fmt"

func main(){
	myNumber  := 23

	// reference to the mynumber -> gives memory address of myNumber
	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	// *ptr -> gives the value of the pointer/the address -> ptr -> mynumber address -> value at that address
	// gives access to the exact value and whatever action is performed,its performed on the actual values

	*ptr = *ptr * 2
	fmt.Println("New value is: ",myNumber)
}