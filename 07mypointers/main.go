package main

import "fmt"

func main() {
	fmt.Println("welcome to class of pointer")

	// var ptr *int
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 66
	var ptr = &myNumber
	fmt.Println("\nvalue of actual pointer is ", ptr)
	fmt.Println("value hold by pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("\nNew value is : ", ptr)
	fmt.Println("New value is : ", *ptr)
}
