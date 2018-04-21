package main

import "fmt"

const (
	A = iota // This is an untyped integer - Initialised to zero.
	B        // Even if you don't pass the value, iota will detect this next following variable and initialise to one.
	C        // Auto-magically this will be initialised to two. (based on A iota)
)

const (
	D = iota // If you define another set of Constants, and call iota, it will be reset to zero!
	E
	F
)

func main() {
	fmt.Println(A) // prints 0
	fmt.Println(B) // prints 1
	fmt.Println(C) // prints 2

	fmt.Println(D) // prints 0
	fmt.Println(E) // prints 1
	fmt.Println(F) // prints 2
}
