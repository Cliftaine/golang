package main

import (
	"fmt"
	"sync"
)

func timesThree(number int, ch chan int) {
	result := number * 3
	fmt.Println(result)
	ch <- result
}

func main() {
	fmt.Println("We are executing a goroutine")
	ch := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	go timesThree(3, ch)
	go timesThree(4, ch2)

	wg.Wait()

	fmt.Println("Printing")

	result := <-ch

	fmt.Printf("The result is: %v\n", result)
	//fmt.Printf("The result2 is: %v\n", result2)
}
