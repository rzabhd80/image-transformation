package commands

import (
	"image"
	"math"
)

var Transformations map[string]Transformation = map[string]Transformation{
	"IdentityTransformation":          IdentityTransformation{},
	"TranslationTransformation":       TranslationTransformation{translationY: 2, translationX: 4},
	"ShearVerticalTransformation":     ShearVerticalTransformation{shear: 2, vertical: 3},
	"ShearHorizontalTransformation":   ShearHorizontalTransformation{shear: 2, displacement: 3},
	"ScalingReflectionTransformation": ScalingReflectionTransformation{scaleY: 2, scaleX: 4},
	"RotationTransformation":          RotationTransformation{Angle: 75},
}

// Transformation interface
type Transformation interface {
	Apply(img image.Image) image.Image
}

// IdentityTransformation represents the identity transformation
type IdentityTransformation struct{}

func (t IdentityTransformation) Apply(img image.Image) image.Image {
	return img // No change for identity transformation
}

// TranslationTransformation represents the translation transformation
type TranslationTransformation struct {
	translationX, translationY float64 // Translation in x and y directions
}

func (t TranslationTransformation) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	translated := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			translated.Set(x+int(t.translationX), y+int(t.translationY), img.At(x, y))
		}
	}
	return translated
}

// ShearVerticalTransformation represents the vertical shear transformation
type ShearVerticalTransformation struct {
	shear, vertical float64 // shear factor and vertical displacement
}

func (displacement ShearVerticalTransformation) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	sheared := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newX := x
			newY := y + int(displacement.shear)*x + int(displacement.vertical)
			if newY >= bounds.Min.Y && newY < bounds.Max.Y {
				sheared.Set(newX, newY, img.At(x, y))
			}
		}
	}
	return sheared
}

// ShearHorizontalTransformation represents the horizontal shear transformation
type ShearHorizontalTransformation struct {
	shear, displacement float64 // shear factor and horizontal displacement
}

func (t ShearHorizontalTransformation) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	sheared := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newX := x + int(t.shear)*y + int(t.displacement)
			newY := y
			if newX >= bounds.Min.X && newX < bounds.Max.X {
				sheared.Set(newX, newY, img.At(x, y))
			}
		}
	}
	return sheared
}

// ScalingReflectionTransformation represents the scaling/reflection transformation
type ScalingReflectionTransformation struct {
	scaleX, scaleY float64 // Scaling factors for x and y directions
}

func (t ScalingReflectionTransformation) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	scaled := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newX := int(float64(x) * t.scaleX)
			newY := int(float64(y) * t.scaleY)
			scaled.Set(newX, newY, img.At(x, y))
		}
	}
	return scaled
}

// RotationTransformation represents the rotation transformation
type RotationTransformation struct {
	Angle float64 // Angle in radians
}

func (t RotationTransformation) Apply(img image.Image) image.Image {
	bounds := img.Bounds()
	rotated := image.NewRGBA(bounds)
	cx := float64(bounds.Dx()) / 2.0
	cy := float64(bounds.Dy()) / 2.0
	cosTheta := math.Cos(t.Angle)
	sinTheta := math.Sin(t.Angle)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newX := int((float64(x)-cx)*cosTheta - (float64(y)-cy)*sinTheta + cx)
			newY := int((float64(x)-cx)*sinTheta + (float64(y)-cy)*cosTheta + cy)
			if newX >= bounds.Min.X && newX < bounds.Max.X && newY >= bounds.Min.Y && newY < bounds.Max.Y {
				rotated.Set(newX, newY, img.At(x, y))
			}
		}
	}
	return rotated
}
