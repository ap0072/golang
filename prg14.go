// if conditions

package main

import "fmt"

func main() {
	a := "happy"
	if a == "Happy" {
		fmt.Println("hello happy")
	} else if a == "appy" {
		fmt.Println("hello happy2")

	} else {
		fmt.Print("all conditions failed")
	}
}
