package main

import "fmt"

func main() {

	x := 1
	var y *int
	y = &x

	fmt.Println("x es una variable, y un apuntador e imprimimos referencia")
	fmt.Println(x, y)

	fmt.Println("x es una variable, y un apuntador e imprimimos a lo que apunta")
	fmt.Println(x, *y)

	x = 2

	fmt.Println(x, y)
	fmt.Println(x, *y)

	*y = 4
	fmt.Println(x, y)
	fmt.Println(x, *y)
}
