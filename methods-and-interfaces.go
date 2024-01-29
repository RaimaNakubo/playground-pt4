package main

import (
	"fmt"
	"time"
)

// MyError describes information about occured error as: time of occuring(When) and error message(What)
type MyError struct {
	When time.Time
	What string
}

func main() {
	//errors
	if err := run(); err != nil { //if an error occured in run() call
		fmt.Println(err) //display information about error
	}
}

// method Error() implements "Errorer" for MyError type;
// it recieves a pointer to an error and prints formatted string with information about error occured
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// function run() returns address of an instance of MyError;
// instace of MyError contains current time and "error message";
// run() is used as example case to test error message
func run() error {
	return &MyError{
		time.Now(),
		"An error occured",
	}
}
