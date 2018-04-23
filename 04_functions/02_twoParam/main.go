package main

import "fmt"

func main() {
	greet("Luis", "Benavides")
}

//func greet(firstName string, secondName string) { // You can pass two parameters or more at the same time.
//	fmt.Printf("Hello %s %s \n", firstName, secondName)
//}

func greet(firstName, lastName string) {  // If they are both the same type you can just define it once.
	fmt.Printf("Hello %s %s. \n", firstName, lastName)
}