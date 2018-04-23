// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23. Find the sum of
// all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func main() {
	var sumMultiples = 0

	for i := 0; i < 1000; i++ { // From all the numbers between 0 and less than 1000 (999)
		if i%3 == 0 || i%5 == 0 { // if i is divisible by 3 or 5
			sumMultiples = sumMultiples + i // add to sumMultiples
		}
	}

	fmt.Println(sumMultiples) // print out the values of sumMultiples
}
