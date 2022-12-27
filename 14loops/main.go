package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in golang")
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}

	value := 0
	for value < 10 {
		if value == 5 {
			value++
			continue
		} else if value == 7 {
			value++
			goto lco
		}
		fmt.Println("value is : ", value)
		value++
	}
lco:
	fmt.Println("value is 7. it is lucky number")
}
