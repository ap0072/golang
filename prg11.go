// Assignment Operators
package main

import "fmt"

func main() {
	var a int = 10
	var b int = 2
	{
		c := a
		d := b
		c += d // applying the addition and assignment operation
		fmt.Printf("The  value of %d += %d is %d \n", a, b, c)
	}
	{
		c := a
		d := b
		c -= d // applying the subtraction and assignment operation
		fmt.Printf("The  value of %d -= %d is %d \n", a, b, c)
	}
	{
		c := a
		d := b
		c *= d // applying the multiplication and assignment operation
		fmt.Printf("The  value of %d *= %d is %d \n", a, b, c)
	}
	{
		c := a
		d := b
		c /= d // applying the division and assignment operation
		fmt.Printf("The  value of %d /= %d is %d \n", a, b, c)
	}
	{
		c := a
		d := b
		c %= d // applying the modulus and assignment operation
		fmt.Printf("The  value of %d %%= %d is %d \n", a, b, c)
	}

}
