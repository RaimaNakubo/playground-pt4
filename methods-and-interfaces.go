/*
Exercise: Readers
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func main() {
	reader.Validate(MyReader{})
}

// method Read() implements io.Reader interface for MyReader type;
// this implementation replaces value of every byte read by Reader with 'A' ASCII character value (which is 65 int)
// and returns number of bytes read by Reader plus nil(success) error message
func (MyReader) Read(b []byte) (int, error) {
	var bytesRead int = 0

	for i := range b {
		b[i] = 'A'
		bytesRead += 1
		//fmt.Println("Current i: ", i)
		//fmt.Println("Current b[i]: ", b[i])
	}

	//fmt.Println("bytes read counter: ", bytesRead)
	//fmt.Println("length of slice b:", len(b))
	//fmt.Println("capacity of slice b:", cap(b))
	//fmt.Println("Content of slice b: ", b)
	return bytesRead, nil // here you can use len(b) instead of bytesRead and get rid of it completely as len(b) returs length of a b slice

	/*
		Reader always should have a slice of bytes as argument, data read by Reader goes into this slice.
			You can manipulate the data read in slice.
		Reader always should have two returning values: number of bytes it have read as integer and error message.
		Also reader can have a reciver as Type(!not instance) of variable. It means that Reader is implemented by this type.
			All instances of readers of that Type calling method Read() would use this specific implementation.
	*/
}
