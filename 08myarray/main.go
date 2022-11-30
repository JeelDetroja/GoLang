package main

import "fmt"

func main() {
	fmt.Println("welcome to array")
	var friutList [4]string

	friutList[0] = "apple"
	friutList[1] = "tomato"
	friutList[2] = "peach"

	fmt.Println("fruit list is : ", friutList)
	fmt.Println("fruit list is : ", len(friutList))
}
