package main

import "fmt"

//down to up when --defer is used

func main() {
	defer fmt.Println("Hello")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("defer in golang :")
	fmt.Println("hi")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
