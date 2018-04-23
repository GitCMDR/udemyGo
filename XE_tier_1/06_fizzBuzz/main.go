// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print "Fizz" instead of the
// number and for the multiples of five print "Buzz". For
// numbers which are multiples of both three and five print
// "FizzBuzz".
package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 { // if i is divisible by 3 or 5 print fizzbuzz
			fmt.Println("FizzBuzz")
		} else if i%5 == 0 {
			fmt.Println("Buzz") // if i is divisible by 5 print buzz
		} else if i%3 == 0 {
			fmt.Println("Fizz") // if i is divisible by 3 print fizz
		} else {
			fmt.Println(i)
		}
	}
}
