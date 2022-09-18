package main

import "fmt"

func main() {

	slice := []string{"Esto", "es", "un", "slice", "bien", "perron"}

	//for tradicional
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range
	for index, element := range slice {
		fmt.Printf("Index = %d and element = %s\n", index, element)
	}

}
