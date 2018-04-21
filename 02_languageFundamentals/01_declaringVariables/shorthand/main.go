package main

import "fmt"

func main() {
	// Declaring with Shorthand ':=', can only be using inside func. It declares, assigns (initialises).
	a := 22
	b := "luis"
	c := 4.20
	d := true

	// %v prints the value in default format, use %T to print the var types.
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
