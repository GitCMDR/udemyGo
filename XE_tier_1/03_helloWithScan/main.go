/* Create a program that prints to the terminal asking for a user to enter
   their name. Use fmt.Scan to read a user’s name entered at the terminal.
   Print “Hello <NAME>” with <NAME> replaced with what the user entered at
   the terminal.*/
package main

import "fmt"

func main() {
	var userName string // initialise variable to store input string

	fmt.Println("Please input your name:") // print instructions to terminal
	fmt.Scanln(&userName)                  // scan input *after new line* and store it in userName variable
	fmt.Printf("Hello %s \n", userName)

}
