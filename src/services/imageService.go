package services

import (
	"image"
	"os"
)

type ImageService struct {
	imagePath string
}

func(this *ImageService) loadImage(imageSource string) (img image.Image, err error){
	file, err := os.Open(imageSource)
	if err != nil {
		return
	}

	defer file.Close()
	img,_, err = image.Decode(file)
	return
}

func(this *ImageService) changeSize(width, height int) {

}