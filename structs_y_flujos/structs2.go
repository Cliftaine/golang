package main

import "fmt"

func main() {

	var person struct {
		name string
		age  int
		pet  string
	}

	fmt.Println(person)

	person.name = "bob"
	person.age = 50
	person.pet = "dog"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)

	//
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}
