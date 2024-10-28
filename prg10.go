// Arthimetic Operators
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2
	fmt.Printf("The  sum of %d + %d is %d \n", a, b, a+b)
	fmt.Printf("The  difference of %d - %d is %d \n", a, b, a-b)
	fmt.Printf("The  multiplication of %d * %d is %d \n", a, b, a*b)
	fmt.Printf("The  division of %d / %d is %d \n", a, b, a/b)
	fmt.Printf("The  modulus of %v %% %v is %d \n", a, b, a%b)
	// we cannot increment the value directly. First we need to do increment operation
	// and then assign to new variable
	c := a
	c++
	d := c
	fmt.Printf("The  value of %d after increment by value one is %d \n", a, d)
	// we cannot increment the value directly. First we need to do increment operation
	// and then assign to new variable
	e := a
	e--
	f := e
	fmt.Printf("The  value of %d after decrement by value one is %d \n", a, f)

}
