package main

import "fmt"

func main() {

	// Crear array
	arr := []string{"Este", "es", "un", "slice"}
	fmt.Println("Array:", arr)

	// Slice a partir de arreglo
	slices := arr[1:4]
	slices2 := make([]string, len(arr))
	slicesCopy2 := copy(slices2, arr)

	slices[0] = "prueba"
	fmt.Println("Array:", arr)
	fmt.Println("Slice1:", slices)
	fmt.Println("Slice2:", slices2)
	fmt.Println("Slicecopiado:", slicesCopy2)
}
