package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice valuse is 1")
	case 2:
		fmt.Println("Dice valuse is 2")
	case 3:
		fmt.Println("Dice valuse is 3")
	case 4:
		fmt.Println("Dice valuse is 4")
		fallthrough
	case 5:
		fmt.Println("Dice valuse is 5")
	case 6:
		fmt.Println("Dice valuse is 6")
	default:
		fmt.Println("what was that")
	}
}
