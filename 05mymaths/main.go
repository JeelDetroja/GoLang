package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths In GoLang..!!")

	//Random Number From - "math/rand"
	// 	rand.Seed(time.Now().UnixNano())
	// 	fmt.Println(rand.Intn(5) + 1)

	//Random Number From - "crypto/rand"
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
