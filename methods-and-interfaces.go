package main

import (
	"fmt"
	"image"
)

func main() {
	//images
	m := image.NewRGBA(image.Rect(0, 0, 100, 100)) //m is instance of RGBA image with boundaries, starting point inside boundaries is 0,0, endpoint is 100,100
	fmt.Println(m.Bounds())                        //printing image's boundaries
	fmt.Println(m.At(0, 0).RGBA())                 //getting color of point of an image at 0,0 coordinates as RGBA value

}
