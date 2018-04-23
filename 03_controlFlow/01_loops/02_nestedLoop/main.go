package main

import "fmt"

// Just like in Python, you can nest for loops within each other.
func main() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 10; j++ {
			fmt.Println(i, " - ", j)
			// End of the (j) nested loop
		}
		// End of the (i) nested loop
	}
}
