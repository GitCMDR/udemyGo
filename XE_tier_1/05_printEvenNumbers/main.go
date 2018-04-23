// Print all of the even numbers between 0 and 100.
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ { // Initialise i to zero, run loop as long as i is less or equal to 100, increment by 1
		if i%2 == 0 { // if remainder of i divided by two (even) is zero
			fmt.Println(i) // print i to terminal
		}
	}
}
