package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels In GoLang..!!")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) { //<-Chan = OutSide Of The Box - Receiving Value

		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { //Chan<- = InSide Of The Box - Sending Value
		myCh <- 7 //Comment THis Line To Get Output False,Chanel Closed.

		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
