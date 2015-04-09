package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 255, 100}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
