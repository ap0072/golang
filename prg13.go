// Bitwise Operators
package main

import "fmt"

func main() {
	var a int = 12
	var b int = 25

	//using Bitwise and operator - &
	c := a & b
	fmt.Printf("the bitwise AND operator value of %d , %d is %d \n", a, b, c)

	//using Bitwise or operator - |
	d := a | b
	fmt.Printf("the bitwise OR operator value of %d , %d is %d \n", a, b, d)

	//using Bitwise xor operator - |
	e := a ^ b
	fmt.Printf("the bitwise XOR operator value of %d , %d is %d \n", a, b, e)

	//using Bitwise Left shift operator - |
	f := 212 << 1
	fmt.Printf("the bitwise Left shift operation on value %d, with position 1 is %d \n", 212, f)

	//using Bitwise Left shift operator - |
	g := 12 >> 2
	fmt.Printf("the bitwise Right shift operation on value %d, with position 2 is %d \n", a, g)

}
