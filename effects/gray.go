package effects

import (
	"github.com/enohr/imgutil-go/effects/helper"
	"image"
	"image/color"
)

func Gray(img image.Image) (grayImage *image.Gray16) {
	xSize := img.Bounds().Size().X
	ySize := img.Bounds().Size().Y
	grayImage = image.NewGray16(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})

	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {
			red, green, blue, _ := img.At(x, y).RGBA()
			redValue := helper.REDWEIGHT * float64(red)
			greenValue := helper.GREENWEIGHT * float64(green)
			blueValue := helper.BLUEWEIGHT * float64(blue)
			sum := redValue + greenValue + blueValue
			newColor := color.Gray16{Y: uint16(sum)}
			grayImage.Set(x, y, newColor)
		}
	}
	return
}
