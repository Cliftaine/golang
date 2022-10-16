package main

import "fmt"

func sum(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Resultado: ", res)
	return 0
}

func show2() {
	fmt.Println("Holas_2")
}

func show() {
	defer show2()
	fmt.Println("Holas")
}

func main() {

	sum(2, 2)

	defer sum(2, 3)
	show()
}
