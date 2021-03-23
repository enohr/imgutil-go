package effects

import (
	"github.com/enohr/imgutil-go/effects/helper"
	"image"
	"math"
)

// Receives an image and radius then returns a new image blurred by gaussian blur
func GaussianBlur(img image.Image, radius int, sigma float64) (blurredImage *image.RGBA) {
	if radius%2 == 0 {
		radius += 1
	}
	var kernel = createGaussianKernel(radius, sigma)

	blurredImage = helper.Convolution(img, kernel, radius)
	return
}

func createGaussianKernel(radius int, sigma float64) (kernel [][]float64) {
	kernel = make([][]float64, radius)
	for i := range kernel {
		kernel[i] = make([]float64, radius)
	}
	var sum float64
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
