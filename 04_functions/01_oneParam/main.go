package main

import "fmt"

func main() { // Call the function greet() and pass a name.
	greet("Luis")
	greet("Zammit")
}

func greet(name string) { // Define function and define parameters (and types) to pass into it.
	fmt.Println(name) // What the function does with 'name' parameter.
}
