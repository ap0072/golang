// Using Scanf function to take the input from user
package main

import "fmt"

func main() {
	var name string
	var phoneNumber int

	fmt.Print("Enter your name and number ")
	//fmt.Scanf("%s %d", &name, &phoneNumber) // Use Scan here
	count, errors := fmt.Scanf("%s %d", &name, &phoneNumber)

	fmt.Printf("Hi, everyone. I am %s and my phone number is %d \n", name, phoneNumber)
	fmt.Printf("Printing the count %v and error %v occur during taking input using Scang function", count, errors)

}
