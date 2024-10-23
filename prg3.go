package main

import "fmt"

func main() {
	var name string
	var phoneNumber int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name) // Use Scan here

	fmt.Print("Enter your phone number: ")
	fmt.Scan(&phoneNumber) // Use Scan here

	fmt.Printf("Hi, everyone. I am %s and my phone number is %d\n", name, phoneNumber)
}
