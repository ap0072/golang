package main

import "fmt"

func two() {

}

func main() {
	var (
		name   string = "prasanth"
		grades int    = 42
	)
	fmt.Printf("hi, everyone. I am %q and I got %d/100 in maths\n ", name, grades)

	{
		var blood_group string = "O+ve"
		fmt.Printf("hi, everyone. I am %q and I got %d/100 in maths and my blood group is %q", name, grades, blood_group)

	}

}
