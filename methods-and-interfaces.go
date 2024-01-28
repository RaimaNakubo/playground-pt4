package main

import (
	"fmt"
	"math"
)

// interface I is a generic interface
type I interface {
	M()
}

// T is a custom type for string
type T struct {
	S string
}

// F is a custom type for float64
type F float64

func main() {
	var i I //creating an instance of interface I

	describe(i)
	//i.M() //"runtime error: invalid memory address or nil pointer dereference" aka "null pointer exeption"
	//there is no method with <nil> type reciever
	fmt.Println()

	var t *T
	i = t //passing nil value with type *T
	describe(i)
	i.M()
	fmt.Println()

	i = &T{"Hello"} //assigning value to the instance i
	describe(i)     //calling a function to see content of i instance's value
	i.M()           //calling a method to see i's value
	fmt.Println()

	i = F(math.Pi) //
	describe(i)    // same thing here
	i.M()          //
	fmt.Println()

}

// method M() recieves a pointer to custom string and prints it's value
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// method M() recieves a custom float64 and prints the value of it's copy
func (f F) M() {
	fmt.Println(f)
}

// function describe() prints value and type of the instance of an interface I
func describe(i I) {
	fmt.Printf("Interface instance's value: %v, type: %T)\n", i, i)
}
