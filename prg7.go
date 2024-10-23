package main

import "fmt"

func main() {
	var name string
	var phoneNumber int

	fmt.Print("Enter your name and number ")
	fmt.Scanf("%s %d", &name, &phoneNumber) // Use Scan here

	fmt.Printf("Hi, everyone. I am %s and my phone number is %d", name, phoneNumber)

	//fmt.Println(count)
	//fmt.Println(error)
}
