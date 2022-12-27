package main

import "fmt"

func main() {
	fmt.Println("if-else in golang")

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exetly 10 count"
	}
	fmt.Println(result)

	if num := 4; num < 10 {
		fmt.Println("lessthan 10")
	} else {
		fmt.Println("not lessthan 10")
	}

	// if err != nil{

	// }
}
