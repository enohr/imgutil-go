package helper

import (
	"image"
	"image/draw"
)

func CopyImage(img image.Image) (newImage *image.RGBA) {
	newImage = image.NewRGBA(img.Bounds())
	draw.Draw(newImage, img.Bounds(), img, img.Bounds().Min, draw.Src)
	return
}
