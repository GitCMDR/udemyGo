package main

import "fmt"

// There is a very cool thing about Go, and that is, you can initialise a
// list of constants if you want using the below syntax.


// Defining constants to pass to main method.
const (
	Pi = 3.14
	Language = "Spanish"
)


func main() {
	// We use Println here. Remember, it includes a '/n' already.
	fmt.Println(Pi)
	fmt.Println(Language)
}



