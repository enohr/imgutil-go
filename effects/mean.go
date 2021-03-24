package effects

import (
	"github.com/enohr/imgutil-go/effects/helper"
	"image"
)

func MeanBlur(img image.Image, radius int) (blurredImage *image.RGBA) {
	if radius%2 == 0 {
		radius += 1
	}
	var kernel = createMeanKernel(radius)

	blurredImage = helper.Convolution(img, kernel, radius)
	return
}

func createMeanKernel(radius int) (kernel [][]float64) {
	kernel = make([][]float64, radius)
	for i := range kernel {
		kernel[i] = make([]float64, radius)
	}
	var sum float64
	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			kernel[i][j] = 1
			sum += 1
		}
	}

	for i := 0; i < len(kernel); i++ {
		for j := 0; j < len(kernel[i]); j++ {
			kernel[i][j] /= sum
		}
	}
	return
}
