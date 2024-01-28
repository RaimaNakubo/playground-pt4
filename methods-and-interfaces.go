package main

import (
	"fmt"
)

// type Person describes a person by it's name and age
type Person struct {
	Name string
	Age  int
}

func main() {
	//type switches
	do(21)      //testing type int
	do("hello") //testing type string
	do(true)    //testing type bool
	fmt.Println()

	//stringers
	person1 := Person{"Roman Muchalov", 25}
	person2 := Person{"Nikolay Dolzhansky", 44}
	fmt.Println(person1, person2)
	//fmt.Printf("%v, %v", person1, person2)
}

// function do() runs a type swtch test for interface{} and prints statements for int, string and unknown types
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)

	}
}

// method String() recieves value Person and returns a formatted string to print;
// this method implements Stringer interface from print.go, bc it has String() method name and type of returning value is string
// if a Stringer is present it redefines how strings are printed inside current file
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
 Stringer is implemented by any value that has a String method,           \fmt\print.go
 which defines the “native” format for that value.
 The String method is used to print values passed as an operand
 to any format that accepts a string or to an unformatted printer
 such as Print.
*/
