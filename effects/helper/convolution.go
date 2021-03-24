package helper

import (
	"image"
	"image/color"
)

func Convolution(img image.Image, kernel [][]float64, radius int) (newImage *image.RGBA) {
	xSize := img.Bounds().Max.X
	ySize := img.Bounds().Max.Y
	newImage = image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})
	for x := img.Bounds().Min.X; x < xSize; x++ {
		for y := img.Bounds().Min.Y; y < ySize; y++ {
			var redValue, greenValue, blueValue float64
			for kernelX := 0; kernelX < len(kernel); kernelX++ {
				for kernelY := 0; kernelY < len(kernel); kernelY++ {
					kernelValue := kernel[kernelX][kernelY]
					red, green, blue, _ := img.At(x+kernelX-radius/2, y+kernelY-radius/2).RGBA()
					redValue += float64(red) * kernelValue
					greenValue += float64(green) * kernelValue
					blueValue += float64(blue) * kernelValue
				}
			}
			newColor := color.RGBA64{R: uint16(redValue), G: uint16(greenValue), B: uint16(blueValue), A: uint16(1)}
			newImage.Set(x, y, newColor)
		}
	}
	return
}
