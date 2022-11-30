package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Rating for our Service : ")

	//coma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)
	fmt.Printf("Type : %T", input)
}
