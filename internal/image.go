package internal

import (
	"image"
	"image/draw"
	"image/png"
	"io"
	"os"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . ImageEncoder
type ImageEncoder func(w io.Writer, i image.Image) error

var EncodeImage ImageEncoder = png.Encode

//counterfeiter:generate . ImageWriter
type ImageWriter func(file string, i image.Image) error

var WriteImage ImageWriter = func(file string, i image.Image) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	err = EncodeImage(f, i)
	if err != nil {
		return err
	}
	return f.Close()
}

//counterfeiter:generate . ImageMaker
type ImageMaker func(width, height int) *image.RGBA

var NewImage ImageMaker = func(width, height int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, width, height))
}

//counterfeiter:generate . Drawer
type Drawer func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point, op draw.Op)

var Draw Drawer = draw.Draw
