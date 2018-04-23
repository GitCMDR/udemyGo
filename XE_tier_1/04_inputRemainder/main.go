// Create a program that prints to the terminal asking for a user
// to enter a small number and a larger number. Print the remainder
// of the larger number divided by the smaller number.

package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int
	var calculatedReminder int

	fmt.Println("Hello! This program will allow you to calculate the remainder of two numbers.")
	fmt.Println("Please input your largest number.")
	fmt.Scanln(&largeNumber)
	fmt.Println("Please input your smallest number.")
	fmt.Scanln(&smallNumber)

	calculatedReminder = largeNumber % smallNumber

	fmt.Printf("The remainder of these two numbers is %v \n", calculatedReminder)
}
