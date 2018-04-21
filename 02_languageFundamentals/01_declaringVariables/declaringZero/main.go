package main

import "fmt"

func main () {
	// We can just declaring the variable and implicitly initialise it to zero value depending on the type
	// eg. int will be automatically be initialised to zero
	// bool will be automatically initialised to false
	// string will be an empty string ('')

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
