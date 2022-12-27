package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var friutList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of friutlist is %T \n", friutList)
	friutList = append(friutList, "mango", "banana")
	fmt.Println(friutList)
	//var friutList [5]string
	friutList = append(friutList[:3])
	fmt.Println(friutList)
	highScores := make([]int, 4)
	highScores[0] = 232
	highScores[1] = 532
	highScores[2] = 26
	highScores[3] = 453
	//highScores[4] = 133

	highScores = append(highScores, 543, 242, 343)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove value from slices based on index

	var courses = []string{"react", "java", "c++", "paython", "ruby", "golang"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
