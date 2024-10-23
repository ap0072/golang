package main

import (
	"fmt"
	"reflect"
)

func main() {

	var name string = "Prasanth"
	//var grades int = 10
	fmt.Printf("hi, everyone. I am %v and type of variable name is %T \n", name, name)
	fmt.Printf("hi, everyone. I am %v and type of variable name is %v", name, reflect.TypeOf(name))

}
