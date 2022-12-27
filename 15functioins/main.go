package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	greeterTwo()
	proRes, message := proAdder(4, 6, 3, 7)
	fmt.Println("Result is: ", proRes, message)
	//fmt.Println(message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "from proResult function"
}
func greeterTwo() {
	fmt.Println("user of pro adder")
}

func greeter() {
	fmt.Println("Hello use of adder")
}
