package effects

import (
	"github.com/enohr/imgutil-go/effects/helper"
	"image"
)

func Gray(img image.Image) (grayImage *image.RGBA) {
	finalX, finalY := img.Bounds().Dx(), img.Bounds().Dy()
	startX, startY := img.Bounds().Min.X, img.Bounds().Min.Y

	copyImage := helper.CopyImage(img)
	grayImage = image.NewRGBA(copyImage.Bounds())

	for x := startX; x < finalX; x++ {
		for y := startY; y < finalY; y++ {
			position := y*copyImage.Stride + x*4

			red := copyImage.Pix[position]
			green := copyImage.Pix[position+1]
			blue := copyImage.Pix[position+2]

			redValue := helper.REDWEIGHT * float64(red)
			greenValue := helper.GREENWEIGHT * float64(green)
			blueValue := helper.BLUEWEIGHT * float64(blue)

			sum := redValue + greenValue + blueValue

			grayImage.Pix[position] = uint8(sum)
			grayImage.Pix[position+1] = uint8(sum)
			grayImage.Pix[position+2] = uint8(sum)
		}
	}
	return
}
