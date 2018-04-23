package main

import "fmt"

// Print Integer to Decimal(Base10) and Binary(Base2) and Hexadecimal(Base16)
func main() {
	// Output of below is 42 - 101010 - 2a
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	// Check out this hexadecimal formatting trick - Print it in CAPITAL letters.
	fmt.Printf("%X \n", 42)
	// Now, if you want the 0X in the beginning you add the #
	fmt.Printf("%#X \n", 42)
}
