package main

import (
	"fmt"
)

func main() {
	var base float32
	var height float32
	const initial float32 = 0.5

	fmt.Println("Enter the base of the triangle: ")
	fmt.Scan(&base) // Use %f for float32

	fmt.Println("Enter the height of the triangle: ")
	fmt.Scan(&height) // Use %f for float32

	formula := initial * base * height
	fmt.Printf("Area of the Triangle is %.2f\n", formula) // Print the area
}
