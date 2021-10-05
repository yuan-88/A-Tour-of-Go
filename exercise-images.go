package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

/* Override methods of Bounds(), ColorModel(), At(x, y int) for Image type */

/* Override Bounds(): 
(original) func (p *RGBA) Bounds() image.Rectangle 
(modified) return image.Rect(0, 0, w int, h int)
*/
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

/* Override ColorModel(): 
(original) func (p *RGBA) ColorModel() color.Model
(modified) return color.RGBAModel  // RGBAModel is a type in the color package
*/
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

/* Override At(x, y int): 
(original) func (p *RGBA) At(x, y int) color.Color
(modified) return color.RGBA{x uint8, y uint8, 255, 255}  // RGBA{} is a type in the color package
*/
func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}