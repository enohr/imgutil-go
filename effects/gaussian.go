package effects

import (
	"image"
	"image/color"
)

// Receives an image and radius then returns a new image blurred by gaussian blur
func GaussianBlur(img image.Image, radius int) (blurredImage *image.RGBA) {
	xSize := img.Bounds().Size().X
	ySize := img.Bounds().Size().Y

	blurredImage = image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})

	var kernel = [][]int{
		{1, 2, 1},
		{2, 4, 2},
		{1, 2, 1},
	}

	// sum := 16

	for x := 1; x < xSize-1; x++ {
		for y := 1; y < ySize-1; y++ {

			var redValue, greenValue, blueValue uint32

			for kernelX := 0; kernelX < 3; kernelX++ {
				for kernelY := 0; kernelY < 3; kernelY++ {
					kernelValue := kernel[kernelX][kernelY]

					red, green, blue, _ := img.At(x-kernelX, y-kernelY).RGBA()
					redValue += red * uint32(kernelValue)
					greenValue += green * uint32(kernelValue)
					blueValue += blue * uint32(kernelValue)
				}
			}
			redValue = redValue / 16
			greenValue = greenValue / 16
			blueValue = blueValue / 16
			newColor := color.RGBA64{R: uint16(redValue), G: uint16(greenValue), B: uint16(blueValue), A: 1}
			blurredImage.Set(x, y, newColor)
		}
	}

	return
}
