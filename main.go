package main

import (
	"image"
	"image/jpeg"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	img := getFileFromPath("image1.jpg")
	saveImage("image2.jpg", img)
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
