package main

import "fmt"

func main() {
	// Statement that is going to be sent to the switch-case loop
	switch "Luis" {
	// Options for the cases and the code that is going to be run.
	case "Luis":
		fmt.Println("Hemlo Master.")
		// 'fallthrough will evaluate the next case to True as well!
		fallthrough
	case "Adam":
		fmt.Println("I don't know who you are.")
	default:
		fmt.Println("I see nobody.")
	}
}
