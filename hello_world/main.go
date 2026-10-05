package main

import "fmt"

// entrypoint of the program
func main(){

	//simple values
	//integers
	fmt.Println(1 + 1)

	// strings11
	fmt.Println("hello word")

	// bool
	fmt.Println(true)
	fmt.Println(false)

	// float
	fmt.Println(10.5)
	fmt.Println(9.0 / 4.0)

	// variables

	// var name string = "golang"
	
	//infer
	var name = "Hello"
	var isValue = true

	fmt.Println(name)
	fmt.Println(isValue)

	// shorthand syntax

	isName := "golang"
	fmt.Println(isName)
}