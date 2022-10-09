package main

import (
	"fmt"
	"time"
)

type person struct {
	name  string
	age   int
	stamp time.Time
}

func main() {

	beth := person{
		age:   30,
		name:  "Beth",
		stamp: time.Now(),
	}

	beth.name = "Bob"
	fmt.Println(beth.stamp)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	for k, v := range teams {
		fmt.Println(k, v)
	}

}
