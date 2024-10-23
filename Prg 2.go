package main

import "fmt"

func main() {
	var (
		name   string
		grades int
	)
	fmt.Printf("hi, everyone. I am %q and I got %d/100 in maths\n ", name, grades)

	{
		var blood_group string = "O+ve"
		fmt.Printf("hi, everyone. I am %q and I got %d/100 in maths and my blood group is %q", name, grades, blood_group)

	}

}
