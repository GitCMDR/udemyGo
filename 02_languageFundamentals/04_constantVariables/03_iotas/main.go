package main

import "fmt"

const (
	A = iota // This is an untyped integer - Initialised to zero.
	B        // Even if you don't pass the value, iota will detect this next following variable and initialise to one.
	C        // Auto-magically this will be initialised to two. (based on A iota)
)

func main() {
	fmt.Println(A) // prints 0
	fmt.Println(B) // prints 1
	fmt.Println(C) // prints 2
}
