package main

import "fmt"

func main() {
	//outer block
	var (
		name   string = "Prasanth"
		grades int    = 75
	)
	fmt.Printf("hi, everyone. I am %v and I got %d/100 in maths\n ", name, grades)

	{
		//inner block
		var blood_group string = "O+ve"
		fmt.Printf("hi, everyone. I am %s and I got %d/100 in maths and my blood group is %q", name, grades, blood_group)

	}

}
