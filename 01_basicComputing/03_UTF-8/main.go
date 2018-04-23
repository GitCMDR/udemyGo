package main

import "fmt"

func main() {
	for i := 0; i < 1888888888; i++ {
		// Print i in decimal, binary, hexadecimal, UTF-8
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
	}

}
