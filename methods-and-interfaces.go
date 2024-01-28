package main

import (
	"fmt"
)

func main() {
	//empty interfaces
	var i interface{} //i is the instance of an empty interface i(nil,nil)
	describe(i)

	i = 42 //i(42,int)
	describe(i)

	i = "hello" //i("hello",string)
	describe(i)

	//fmt.Print() <-- takes any number of arguments of interface{} type
	fmt.Println()

	//interface's type assertions
	var j interface{} = "hello" //j("hello", string)

	s := j.(string) //s("hello",string)
	fmt.Println(s)

	s, ok := j.(string) //s == hello, ok == true
	fmt.Println(s, ok)

	f, ok := j.(float64) //f == 0.0, ok == false
	fmt.Println(f, ok)

	/*
		f = j.(float64) //panic "interface conversion: interface {} is string, not float64"
						//this happened bc statement j.(float64) asserts that value j holds float64 type, but actually it holds string type; do not test like this
		fmt.Println(f)
	*/
}

// function describe() prints value and type of interface{} instance
func describe(i interface{}) {
	fmt.Printf("Interface instance value: %v, type: %T\n", i, i)
}
