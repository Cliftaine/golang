package main

import "fmt"

var (
	a = 654
	b = false
	c = 2.651
	d = 4 + 1i
	e = "Australia"
	f = 15.2 * 4525.321
)

func main() {

	fmt.Printf("s for string: %s\n", variable)
	fmt.Printf("d for Integer: %d\n", a)
	fmt.Printf("6d for Integer: %6d\n", a)

	fmt.Printf("t for Boolean: %t\n", b)
	fmt.Printf("g for Float: %g\n", c)
	fmt.Printf("e for Scientific Notation: %e\n", f)
	fmt.Printf("E for Scientific Notation: %E\n", f)
	fmt.Printf("s for String: %s\n", e)
	fmt.Printf("G for Complex: %G\n", d)

	fmt.Printf("15s String: %15s\n", e)
	fmt.Printf("-10s String: %-10s\n", e)

	t := fmt.Sprintf("Print from right: %[3]d %[2]d %[1]d\n", 11, 22, 33)
	fmt.Println(t)
}
