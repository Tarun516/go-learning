package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop but no keyword

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// infinite loop 
	// for {
		// println("1")
	// }

	// classic for loop
	// for i := 0; i < 3; i++ {

	// 	//break -> break the loop (stop)
	// 	// continue -> skip the current iteration

	// 	fmt.Println(i)
	// }

	// range
	// for i := range 3{
		// fmt.Println(i)
	// }
}