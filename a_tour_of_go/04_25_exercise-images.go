package main

import "golang.org/x/tour/pic"

import "image"
import "image/color"

type Image struct{}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 46, 49)
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{127, 127, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
