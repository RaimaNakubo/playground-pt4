package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	//Readers
	r := strings.NewReader("Hello, Reader!") //r is a reader, reading from string "Hello, Reader!"

	b := make([]byte, 8) //b is a slice of bytes with length and capacity of 8

	for { //infinite loop start
		n, err := r.Read(b) //reading from a reader r to a slice b until slice is full, returning number of bytes read in n and error message in err
		fmt.Printf("Bytes read: %v, error: %v, read bytes in slice: %v\n", n, err, b)
		fmt.Printf("content of slice: %q\n", b[:n])

		if err == io.EOF { //if the end of the string reached
			break //loop ends
		}
	}
}
