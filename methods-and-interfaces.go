/*
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

Create a new type: type ErrNegativeSqrt float64

and make it an error by giving it a: func (e ErrNegativeSqrt) Error() string

method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/

package main

import (
	"fmt"
)

// ErrNegativeSqrt describes information about error occured when caclulating square root as: negative number argument
type ErrNegativeSqrt float64

func main() {

	//fmt.Println(ErrNegativeSqrt(-2).Error()) //errorer test
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}

// function Sqrt() calculates square root of float64 and returns float64 result + error message
func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(x)
	buffer := float64(0)
	change := float64(1)

	fmt.Println("Calculating square root of", x)
	for i := 0; i < 10; i++ {
		buffer = z
		z -= (z*z - x) / (2 * z)
		change = buffer - z
		fmt.Println("Iteration", i, ":", z, "change:", change)

		if change <= 1e-16 {
			fmt.Printf("Calculation completed in %v iterations\n", i+1)
			break
		}
	}

	return z, nil
}

// method Error() implements "Errorer" for ErrNegativeSqrt type;
// error output: message, negative argument value
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))

	/*
		Returning value of this method is a string with error message and a value of a variable that caused error.
		When error message is just a string, value that you're looking for is stored in instance e of ErrNegativeSqrt type.

		If you'd pass e just like it is: return fmt.Sprintf("Cannot Sqrt negative number: %v", e),
			it would cause runtime error - Exception has occurred: fatal error "stack overflow"

		This would happen because  - fmt.Sprintf format %v with arg e causes recursive (playground-pt4.ErrNegativeSqrt).Error method callprintf(default)
		which means Error() method call happens over and over until the content of call stack in your machine reaches boundaries of a call stack itself,
			resulting call stack overflow.
		Recursive call of a Error() method happens because instance of e is passed as argument in fmt.Sprintf()
		Instance of e has ErrNegativeSqrt type (which is custom type for error in this file) that implements "Errorer" interface in this Error() method,
			so when any print call happens (fmt.Sprintf() in this case) an Error() method is called again resulting another call of fmt.Sprintf(),
			which leads to an infinite loop.

		To overcome this behaviour, you should cast instance e type to any other type that is not an error type.
		For this implementation error type is ErrNegativeSqrt, so if you cast the type of e to float64 for example, no infinite loop would occur and Error()
		method would return correct value.

	*/
}
