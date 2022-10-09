package main

import (
	"fmt"
	"sort"
)

func main() {

	scl1 := []int{400, 600, 100, 300, 500, 200, 900}
	scl2 := []int{-23, 567, -34, 67, 0, 12, -5}

	// Displaying slices
	fmt.Println("Slices(Before):")
	fmt.Println("Slice 1: ", scl1)
	fmt.Println("Slice 2: ", scl2)

	// Sorting the slice of ints
	// Using Ints function
	sort.Ints(scl1)
	sort.Ints(scl2)

	// Displaying the result
	fmt.Println("\nSlices(After):")
	fmt.Println("Slice 1 : ", scl1)
	fmt.Println("Slice 2 : ", scl2)
}
