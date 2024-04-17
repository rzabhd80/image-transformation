package encoder

import (
	"errors"
	jpeg "image/jpeg"
	png "image/png"
	"image_transformation/imageLoader"
	"os"
)

func EncodeAndWriteData(img imageLoader.Image, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	if img.Format == "png" {
		return png.Encode(file, img.Data)
	} else if img.Format == "jpeg" {
		return jpeg.Encode(file, img.Data, nil)
	}
	return errors.New("Unsupported image format")
}
