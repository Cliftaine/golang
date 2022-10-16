package main

import "fmt"

func sum(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Resultado: ", res)
	return 0
}

func show() {
	fmt.Println("Holas")
}

func main() {

	sum(2, 2)

	defer sum(5, 5)

	show()
}
