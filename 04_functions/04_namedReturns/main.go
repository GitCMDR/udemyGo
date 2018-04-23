package main

import "fmt"

func main() {
	s := greet("Luis ", "Benavides")
	fmt.Println(s)
}

// You can name what can be returned and return without specifying what at the end.
func greet(firstName, lastName string) (s string) { // You can name what can be returned and return
	s = fmt.Sprint(firstName, lastName)
	return
}
