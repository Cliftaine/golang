package main

import (
	"fmt"
)

func main() {

	abs := map[int]int{
		1: 1,
	}

	fmt.Println(abs[2])

	//maps with arrays
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	for k, v := range teams {
		fmt.Println(k, v)
	}

	fmt.Println(teams)

	//maps with returning
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)

	//DELETE
	n := map[string]int{
		"hello": 5,
		"world": 10,
	}

	fmt.Println(n)

	delete(n, "hello")

	fmt.Println(n)

	//Maps as sets
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

	if intSet[5] {
		fmt.Println("5 is in the set")
	}
}
