package main

import "fmt"

// It will multiply this string by two, a la 'Luis' -> 'Luis Luis'
func main() {
	first, last := greet("Luis ", "Benavides ")
	fmt.Printf("%s %s \n", first, last)
}

// You can return two things in a function, isn't that amazing!
func greet(firstName, lastName string) (string, string) {
	return fmt.Sprint(firstName, firstName), fmt.Sprint(lastName, lastName)
}
