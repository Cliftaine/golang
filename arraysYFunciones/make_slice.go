package main

import "fmt"

func main() {

	//Se crea un arreglo de 7 posiciones y se genera un slice de 4 a partir del mismo
	var slice = make([]int, 5, 8)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		slice, len(slice), cap(slice))

	// Se crea un arreglo de 7 posiciones y se genera una referencia al mismo
	var slice2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		slice2, len(slice2), cap(slice2))

}
