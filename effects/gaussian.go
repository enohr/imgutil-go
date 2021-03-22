package effects

import (
	"image"
	"image/color"
	"math"
)

// Receives an image and radius then returns a new image blurred by gaussian blur
func GaussianBlur(img image.Image, radius int) (blurredImage *image.RGBA) {
	xSize := img.Bounds().Size().X
	ySize := img.Bounds().Size().Y

	blurredImage = image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})

	var kernel = createGaussianKernel(5, 1.5)

	// sum := 16

	for x := 1; x < xSize-1; x++ {
		for y := 1; y < ySize-1; y++ {

			var redValue, greenValue, blueValue, alphaValue uint32

			for kernelX := 0; kernelX < radius; kernelX++ {
				for kernelY := 0; kernelY < radius; kernelY++ {
					kernelValue := kernel[kernelX][kernelY]

					red, green, blue, alpha := img.At(x-kernelX, y-kernelY).RGBA()
					redValue += red * uint32(kernelValue)
					greenValue += green * uint32(kernelValue)
					blueValue += blue * uint32(kernelValue)
					alphaValue += alpha
				}
			}
			alphaValue = alphaValue / 3
			redValue = redValue
			greenValue = greenValue
			blueValue = blueValue
			newColor := color.RGBA64{R: uint16(redValue), G: uint16(greenValue), B: uint16(blueValue), A: uint16(alphaValue)}
			blurredImage.Set(x, y, newColor)
		}
	}

	return
}

func createGaussianKernel(radius int, sigma float64) [5][5]float64 {
	var kernel [5][5]float64
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

	return kernel
}

func calculateGaussFunction(x, y int, sigma float64) (result float64) {
	numerator := -(math.Pow(float64(x), 2) + math.Pow(float64(y), 2))
	denominator := 2 * math.Pow(sigma, 2)
	result = (1 / (2 * math.Pi * math.Pow(sigma, 2))) * math.Exp(numerator/denominator)
	return
}
