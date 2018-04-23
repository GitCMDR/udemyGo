package main

import "fmt"

func main() {
	// 1. Implicitly initialise the variable 'i' with a value of zero
	// 2. As long as 'i' is less than zero, run this code.
	// 3. When core is run, increment 'i' by 1 with ++
	for i := 0; i < 500001; i++ {
		// Print the value of i in %d (decimal), tab, %b(binary), tab, %x(hexadecimal), new line
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
