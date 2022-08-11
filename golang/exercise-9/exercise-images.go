package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

const (
	width  = 255
	height = 255
)

type Image struct{}

func (Image) ColorModel() color.Model { return color.RGBAModel }
func (Image) Bounds() image.Rectangle { return image.Rect(0, 0, width, height) }
func (Image) At(x, y int) color.Color {
	v := uint8(x*x + y*y)
	return color.RGBA{v, v, 255, 255}
}
func main() {
	m := Image{}
	pic.ShowImage(m)
}
