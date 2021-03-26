package effects

import (
	"github.com/enohr/imgutil-go/effects/helper"
	"image"
	"image/color"
	"math"
)

func SobelFilter(img image.Image) (newImage *image.Gray16) {
	radius := 3
	kernelX, kernelY := createKernels()

	xSize := img.Bounds().Max.X
	ySize := img.Bounds().Max.Y
	newImage = image.NewGray16(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})
	for x := img.Bounds().Min.X; x < xSize; x++ {
		for y := img.Bounds().Min.Y; y < ySize; y++ {
			var redValueX, greenValueX, blueValueX float64
			var redValueY, greenValueY, blueValueY float64
			for kX := 0; kX < 3; kX++ {
				for kY := 0; kY < 3; kY++ {
					kernelValueX := kernelX[kX][kY]
					kernelValueY := kernelY[kX][kY]
					red, green, blue, _ := img.At(x+kX-radius/2, y+kY-radius/2).RGBA()

					redValueX += kernelValueX * float64(red)
					greenValueX += kernelValueX * float64(green)
					blueValueX += kernelValueX * float64(blue)

					redValueY += kernelValueY * float64(red)
					greenValueY += kernelValueY * float64(green)
					blueValueY += kernelValueY * float64(blue)

				}
			}
			redValue := math.Sqrt(math.Pow(redValueX, 2) + math.Pow(redValueY, 2))
			greenValue := math.Sqrt(math.Pow(greenValueX, 2) + math.Pow(greenValueY, 2))
			blueValue := math.Sqrt(math.Pow(blueValueX, 2) + math.Pow(blueValueY, 2))

			red := helper.REDWEIGHT * redValue
			green := helper.GREENWEIGHT * greenValue
			blue := helper.BLUEWEIGHT * blueValue
			sum := red + green + blue

			newColor := color.Gray16{Y: uint16(sum)}
			newImage.Set(x, y, newColor)
		}
	}
	return

}

func createKernels() (kernelX, kernelY [][]float64) {
	kernelX = [][]float64{
		{1, 0, -1},
		{2, 0, -2},
		{1, 0, -1},
	}

	kernelY = [][]float64{
		{1, 2, 1},
		{0, 0, 0},
		{-1, -2, -1},
	}
	return
}
