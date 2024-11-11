//  Bitwise Operators
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2

	//using Bitwise and operator - &
	c := a & b
	fmt.Printf("the bitwise AND operator value of %d ,  %d is %d \n", a, b, c)

	//using Bitwise or operator - |
	c := a | b
	fmt.Printf("the bitwise OR operator value of %d ,  %d is %d \n", a, b, c)

	//using Bitwise xor operator - |
	c := a ^ b
	fmt.Printf("the bitwise XOR operator value of %d ,  %d is %d \n", a, b, c)
}
