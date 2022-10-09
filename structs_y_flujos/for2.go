package main

import (
	"fmt"
)

func main() {
	i := 1
	for i < 100 {
		fmt.Println(i)
		a := 100
		fmt.Println(a)
		i = i * 2
	}
}
