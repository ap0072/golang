// Comparison Operators i.e ==, !=, <, <=, >, >=
package main

import "fmt"

func main() {
	var name string = "India"
	var name_2 string = "USA"
	var a, b, c int = 1, 2, 3
	fmt.Printf("The boolean value of %d == %d is %t \n", a, b, a == b)
	fmt.Printf("The boolean value of %d != %d is %t \n", a, b, a != b)
	fmt.Printf("The boolean value of %d < %d is %t \n", a, b, a < b)
	fmt.Printf("The boolean value of %d <= %d is %t \n", a, b, a <= b)
	fmt.Printf("The boolean value of %v == %v is %t \n", name, name_2, name == name_2)
	fmt.Printf("The boolean value of %v != %v is %t \n", name, name_2, name < name_2)
	fmt.Printf("The boolean value of %d > %d is %t \n", b, c, b > c)
	fmt.Printf("The boolean value of %d >= %d is %t \n", b, c, b >= c)

}
