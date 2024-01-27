package main

import (
	"fmt"
	"math"
)

// Vertex is a custom data type
type Vertex struct {
	X, Y float64
}

// custom type for float64
type MyFloat float64

// method Abs() recieves a Vertex and returns float64 copy
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// here Abs() is just a function; it has no reciever, but works exactly the same
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method Abs() recieves a custom Float64 and returns float64 absolute value
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// method pAbs( recieves a point to Vertex and returns float64 absolute value)
func (v *Vertex) pAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// method Scale() recieves a point to Vertex, float64 as argument and returns nothing
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// here Scale() is just a function; it has no reciever, but works exactly the same
func Scale(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) //calling a method for a v Vertex(struct) reciever of a method
	fmt.Println()

	fmt.Println(Abs(v)) //calling a function regularly and passing v Vertex argument in
	fmt.Println()

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) //calling a method for a f MyFloat(no-struct) reciever of a method
	fmt.Println()

	v.Scale(10)          //calling a method for a v *Vertex reciever of a method, changing the original v Vertex
	fmt.Println(v.Abs()) //at this point v is Vertex{30, 40}
	fmt.Println()

	Scale(&v, 10) //calling a function regularly and passing address of v Vertex argument in
	fmt.Println(Abs(v))
	fmt.Println()

	//Methods and pointer indirection
	value := Vertex{3, 4.5}
	value.Scale(2)     //calling a method for actual value(not for value's adress); Go's interpretetion would be (&value).Scale(2)
	Scale(&value, 10)  //calling a function passing address of the actual value as argument
	fmt.Println(value) //both calls works completely fine resulting: value Vertex {60, 90}
	fmt.Println()

	pointer := &Vertex{0.2, 1}
	pointer.Scale(2)      //calling a method for value's address
	Scale(pointer, 10)    //calling a function passing actual pointer to the value as argument
	fmt.Println(pointer)  //both calls works completely fine resulting: address of a Vertex{4, 20} somewhere in the memory
	fmt.Println(*pointer) //adding * results value(Vertex{4, 20}) than lies in a pointer *Vertex
	fmt.Println(&pointer) //adding & results memory address of a Vertex{4, 20}
	fmt.Println()

	value = Vertex{3, 4}
	fmt.Println(value.Abs()) //calling a method that expects to recieve value and recieves value
	fmt.Println(Abs(value))  //calling a function that expects to pass value as argument and passing the value as argument
	fmt.Println()            //regular behaviour above

	pointer = &Vertex{4, 3}
	fmt.Println(pointer.Abs()) //calling a method that expects to recieve value and recieves pointer to a value; Go's interpretation would be (*pointer).Abs()
	fmt.Println(Abs(*pointer)) //calling a function that expects to pass value as argument and passing a pointer to a pointer to a value; compiler adds * automaticly
	fmt.Println()              //irregular behaviour above

	//choosing a value or a pointer reciever
	pointer = &Vertex{3, 4}
	fmt.Printf("Vertex before scaling: %+v, Absolute value: %v\n", pointer, pointer.pAbs()) //method recieves a point to Vertex even if it does not need to change the reciever
	pointer.Scale(3)                                                                        //method recieves a point to Vertex and changes values of a Vertex
	fmt.Printf("Vector after scaling: %+v, Absolute value: %v\n", pointer, pointer.pAbs())  //method doesn't need to change the reciever again but still recieves a pointer to Vertex
}
