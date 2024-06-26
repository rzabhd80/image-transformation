package imageLoader

import (
	"image"
	"image_transformation/errors"
	"os"
)

type ImageLoader interface {
	loadImage(filename string) (image.Image, error)
}

type Image struct {
	Path   string
	Format string
	Data   image.Image
}

func (i Image) loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func NewImage(format string, path string) (*Image, error) {
	var validatedFormat bool = validateFormat(format)
	if !validatedFormat {
		return nil, errors.UnsupportedFormat{}
	}

	var image Image = Image{Format: format, Path: path, Data: nil}
	imageData, err := image.loadImage(path)
	image.Data = imageData
	if err != nil {
		return nil, err
	}
	return &image, nil
}
