package main

import (
	"fmt"
	"github.com/enohr/imgutil-go/effects"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	img := getFileFromPath("image.jpg")
	newImg := effects.GaussianBlur(img, 7, 11)
	saveImage("blurredImage.jpg", newImg)
}

func getFileFromPath(path string) (img image.Image) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func saveImage(name string, img image.Image) {
	file, _ := os.Create(name)
	_ = jpeg.Encode(file, img, nil)
}
