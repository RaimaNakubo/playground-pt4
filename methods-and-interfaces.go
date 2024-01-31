package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// MyImage is a custom type for image interface;
// this type will implement Image interface;
// resulting image will be 255x255 rectangle with color in RGBA model
type MyImage struct {
	//here migh've been some properties for an image
	//like height, width and base colours in RGBA
}

func main() {
	//images
	m := MyImage{}
	pic.ShowImage(&m)

}

// method ColorModel() returns RGBA color model for MyImage type of instance
func (MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

// method Bounds() sets image's boundaries for MyImage type of instance
func (MyImage) Bounds() image.Rectangle {

	return image.Rect(0, 0, 255, 255)
}

// method At() returns the color of an Image point in instance i as RGBA
func (i *MyImage) At(x, y int) color.Color {
	v := uint8((x^y)*((x+y)/2) + (x * y))
	return color.RGBA{v, v, 255, 255}
}
