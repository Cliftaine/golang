package main

import (
	"fmt"
)

func main() {

	evenVals := []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}

	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}
