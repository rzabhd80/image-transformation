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
	format string
	path   string
	data   image.Image
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

func newImage(format string, path string) (*Image, error) {
	var validated_format bool = validateFormat(format)
	if !validated_format {
		return nil, errors.UnsupportedFormat{}
	}

	var image Image = Image{format: format, path: path, data: nil}
	image_data, err := image.loadImage(path)
	image.data = image_data
	if err != nil {
		return nil, err
	}
	return &image, nil
}
