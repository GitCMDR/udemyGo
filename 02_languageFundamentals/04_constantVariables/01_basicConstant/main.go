package _1_basicConstant

import "fmt"

const p string = "Luis"

func main() {

	const q = 50

	fmt.Printf("The value of this p is %s \n", p)
	fmt.Printf("The value of q is %v \n", q)

	// Go blocks you from updating Constant variables. Big difference from Python
	//q = 15
	//fmt.Printf("The new value of q is %v \n", q)
}
