package main

import "fmt"

// A variadic function looks like '...string' or '...int' - A function with such parameter
// is called variadic and may be invoked with zero or more arguments for that parameter.

func main() {
	n := average(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(n)
}

// Function average is a variadic function and takes as many float64 as passed
// to it. It then sums all the floats and divides the total by the element count
// of the elements passed to the function. Returns average.
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
