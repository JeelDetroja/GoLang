package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 Monday"))

	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}
