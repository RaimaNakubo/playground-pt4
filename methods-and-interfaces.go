/*

Exercise: rot13Reader
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.
*/

package main

import (
	"io"
	"os"
	"strings"
)

// rot13Reader is a wrapper for Reader which implements Read() method
type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!") //s is a Reader instace initialized by rot13-ciphered string
	r := rot13Reader{s}                             //wrapping Reader in rot13 Reader, so r is an instance of rot13 Reader { Reader }
	io.Copy(os.Stdout, &r)                          //copying data from rot13 Reader to standert outut channel until EOF or error

}

// method Read() implements Reader for rot13Reader type;
// it recieves a pointer to rot13Reader instance of io.Reader instance,
// deciphers bytes from io.Reader instance through []byte argument,
// and returns number of bytes read from io.Reader instance plus error message
func (r13r *rot13Reader) Read(b []byte) (int, error) {

	inputAlphabet := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	outputAlphabet := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")

	n, err := r13r.r.Read(b)

	//I know that nested "for" loops is really bad idea
	//I should've used map or inline calculation (like this: b[i] = (v-'a'+13)%26 + 'a') instead
	//but idc now, it was my firs idea
	for i, v := range b[:n] {
		for j, vc := range inputAlphabet {
			if vc == v {
				b[i] = outputAlphabet[j]
			}
		}
	}

	return n, err

	/*
		rot13 is a letter substitution cipher.
		Every character in rot13 string has to be replaced with the 13th letter after it in the latin alphabet.
		Because there is 26 (2*13) letters in the basic latin alphabet, rot13 is it's own inverse.
		That means, to undo rot13 cypher the same algorythm has to be applied.
	*/
}
