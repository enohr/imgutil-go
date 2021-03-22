package effects

import (
	"image"
	"image/color"
	"math"
)

// Receives an image and radius then returns a new image blurred by gaussian blur
func GaussianBlur(img image.Image, radius int, sigma float64) (blurredImage *image.RGBA) {
	xSize := img.Bounds().Size().X
	ySize := img.Bounds().Size().Y

	if radius%2 == 0 {
		radius += 1
	}

	blurredImage = image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})

	var kernel = createGaussianKernel(radius, sigma)

	for x := 0; x < xSize; x++ {
		for y := 0; y < ySize; y++ {

			var redValue, greenValue, blueValue float64

			for kernelX := 0; kernelX < radius; kernelX++ {
				for kernelY := 0; kernelY < radius; kernelY++ {
					kernelValue := kernel[kernelX][kernelY]
					red, green, blue, _ := img.At(x+kernelX-radius/2, y+kernelY-radius/2).RGBA()
					redValue += float64(red) * kernelValue
					greenValue += float64(green) * kernelValue
					blueValue += float64(blue) * kernelValue

				}
			}
			newColor := color.RGBA64{R: uint16(redValue), G: uint16(greenValue), B: uint16(blueValue), A: uint16(1)}
			blurredImage.Set(x, y, newColor)
		}
	}

	return
}

func createGaussianKernel(radius int, sigma float64) (kernel [][]float64) {
	kernel = make([][]float64, radius)
	for i := range kernel {
		kernel[i] = make([]float64, radius)
	}
	var sum float64
	var normalized float64
	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			kernelValue := calculateGaussFunction(i-radius/2, j-radius/2, sigma)
			kernel[i][j] = kernelValue
			sum += kernelValue
		}
	}

	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			kernel[i][j] /= sum
			normalized += kernel[i][j]
		}
	}
	return
}

func calculateGaussFunction(x, y int, sigma float64) (result float64) {
	numerator := -(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
	denominator := 2 * math.Pow(sigma, 2)
	result = (1 / (2 * math.Pi * math.Pow(sigma, 2))) * math.Exp(numerator/denominator)
	return
}
