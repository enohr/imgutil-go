package helper

import (
	"image"
	"image/draw"
)

func Convolution(img image.Image, kernel [][]float64, radius int) (newImage *image.RGBA) {
	xSize := img.Bounds().Max.X
	ySize := img.Bounds().Max.Y

	copyImage := image.NewRGBA(img.Bounds())
	draw.Draw(copyImage, img.Bounds(), img, img.Bounds().Min, draw.Src)

	newImage = image.NewRGBA(image.Rectangle{Min: image.Point{}, Max: image.Point{X: xSize, Y: ySize}})

	for x := img.Bounds().Min.X + radius/2; x < xSize-radius/2; x++ {
		for y := img.Bounds().Min.Y + radius/2; y < ySize-radius/2; y++ {
			var redValue, greenValue, blueValue float64
			for kernelX := 0; kernelX < len(kernel); kernelX++ {
				for kernelY := 0; kernelY < len(kernel); kernelY++ {
					kernelValue := kernel[kernelX][kernelY]
					xPosition := x + kernelX - radius/2
					yPosition := y + kernelY - radius/2

					position := yPosition*copyImage.Stride + xPosition*4

					red := copyImage.Pix[position]
					green := copyImage.Pix[position+1]
					blue := copyImage.Pix[position+2]

					redValue += float64(red) * kernelValue
					greenValue += float64(green) * kernelValue
					blueValue += float64(blue) * kernelValue
				}
			}
			pixPosition := (y-radius/2)*newImage.Stride + (x-radius/2)*4
			newImage.Pix[pixPosition] = uint8(redValue)
			newImage.Pix[pixPosition+1] = uint8(greenValue)
			newImage.Pix[pixPosition+2] = uint8(blueValue)
			newImage.Pix[pixPosition+3] = 255
			// Use pix slice instead of Set/At is more efficient.
		}
	}
	return
}
