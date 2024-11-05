// Logical Operators
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2
	fmt.Printf("the boolean value of %d > %d , is %t \n", a, b, a > b)
	fmt.Printf("the boolean value of %d < %d , is %t \n", a, b, a < b)
	//using Logical and operator - &&
	c := (a > b) && (a < b)
	fmt.Printf("the boolean value of (%d > %d) and (%d < %d) , is %t \n", a, b, a, b, c)
	//using Logical or operator - ||
	d := (a > b) || (a < b)
	fmt.Printf("the boolean value of (%d > %d) or (%d < %d) , is %t \n", a, b, a, b, d)
	//using Logical not operator - !
	fmt.Printf("the boolean value of not of false , is %t \n", !c)
	fmt.Printf("the boolean value of not of true , is %t \n", !d)
}
