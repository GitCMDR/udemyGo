package main

import "fmt"

func main() {
	s := greet("Luis ", "Benavides")
	fmt.Println(s)
}

func greet(firstName, lastName string) string {
	return fmt.Sprint(firstName, lastName)  // Sprint concatenates strings
}
