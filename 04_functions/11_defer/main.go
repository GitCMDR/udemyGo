package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Print("World \n")
}

// This should print "world hello" - You can defer a call until before
// the execution of a program. And you'll get "Hello World"
func main() {
	defer world()
	hello()
}
