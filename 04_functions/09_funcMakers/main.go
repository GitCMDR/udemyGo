package main

import "fmt"

// Can also do something similar to a function type initialiser!
func makeGreeter() func() string {
	return func() string {
		return "Hello World"
	}
}

func main() {
	greetR := makeGreeter()
	fmt.Println(greetR())

}
