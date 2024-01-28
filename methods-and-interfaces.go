package main

import (
	"fmt"
	"math"
)

// interface Abser holds signature for methods which calculate absolute values
type Abser interface {
	Abs() float64
}

// interface I is declared implicitly, it holds signature of some generic method M(); returning value is not present
type I interface {
	M()
}

// custom type for string
type T struct {
	S string
}

// custom type for float64
type MyFloat float64

// Vertex{X,Y} type
type Vertex struct {
	X, Y float64
}

func main() {
	var a Abser               //a is the instance of Abser interface
	f := MyFloat(-math.Sqrt2) //f is the instance of a custom float64
	v := Vertex{3, 4}         //v is the instance of a Vertex

	a = f //f custom float64 implements Abs() method for a Abser, so a MyFloat implements Abser interface
	fmt.Println(a.Abs())

	a = &v //v *Vertex implements Abs() method for a Abser, so a *Vertex implements Abser interface
	fmt.Println(a.Abs())

	/*
		a = v //cannot use v (variable of type Vertex) as Abser value in assignment: Vertex does not implement Abser (method Abs has pointer receiver)compilerInvalidIfaceAssign)
		fmt.Println(a.Abs())

		There is a compile error on the line 33.  v Vertex cannot implement Abser interface bc v Vertex is not the type of any existing recievers.
		Method Abs() recieves eiter custom float64 or pointer to a Vertex, and v is a Vertex type, not *Vertex nor custom float64, so v Vertex cannot
		be passed as a reciever, therefore cannot implement Abser interface.
	*/
	fmt.Println()

	var i I = T{"hello"} //implementing implicit interface I with custom string of T type; instance i of interface I is no longer implicit interface
	i.M()                //so method M() is no longer generic bc M() implementation for i explicit interface is present
}

// method Abs() recieves custom float64 value, calculates absolute value of the reciever and returns result as float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// method Abs() recieves point to Vertex, calculates absolure value of a Vertex and returns result as float64
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method M() recieves some "value" T and prints the result of a method S() call from copy of a recieved value
// type T implements interface I via adding some specifics to the generic method M(), reciever type in this case; no explicit declaration of implementation needed
func (t T) M() {
	fmt.Println(t.S)
}
