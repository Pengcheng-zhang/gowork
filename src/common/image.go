package common

import (
	"image"
	"os"
)

type ImageBiz struct {
	imagePath string
}

func(this *ImageBiz) loadImage(imageSource string) (img image.Image, err error){
	file, err := os.Open(imageSource)
	if err != nil {
		return
	}

	defer file.Close()
	img,_, err = image.Decode(file)
	return
}

func(this *ImageBiz) changeSize(width, height int) {

}