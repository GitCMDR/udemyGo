package main

import "fmt"

// You would normally call a function by first defining it outside main and then calling it from main.
//func main() {
//	greeting()
//}
//
//func greeting() {
//	fmt.Println("Hello!")
//}

// But what if you wanted to define a function within main? Well, for that you
// need Func expressions! It is the only way you can have a function within
// another function.

func main() {
	greeting := func() {
		fmt.Println("Hello!")
	}

	greeting()
}
