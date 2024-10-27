package main

import (
	"fmt"
	"strconv"
)

func main() {

	var strname string = "200"
	var intgrades int = 10

	gradesinstring := strconv.Itoa(intgrades) // converting the integer variable to String variable
	fmt.Printf("hi, everyone grades are now printing in string format %v and type of this variable intgrades is now %T \n", intgrades, gradesinstring)

	nameininteger, errors := strconv.Atoi(strname)
	fmt.Printf("hi, everyone. my string variable  is now converted into integer variable with value %v %T \n", nameininteger, nameininteger)
	fmt.Printf("%v %T", errors, errors)

}
