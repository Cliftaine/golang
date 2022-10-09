package main

import "fmt"

func main() {

	// Definición del arreglo 1D
	arr := []string{"Arreglo", "de", "strings"}

	fmt.Println("Elementos del arreglo")

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//arreglo 2D
	arr2d := [][]string{
		{"C #", "C", "Python"},
		{"Java", "Scala"},
		{"C++"}}
	fmt.Println("Elements of Array 1")
	for x := 0; x < len(arr2d); x++ {
		fmt.Println("------")
		for y := 0; y < len(arr2d[x]); y++ {
			fmt.Println(arr2d[x][y])
		}
	}
}
