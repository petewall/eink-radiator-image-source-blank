package lib

import (
	"image/color"

	"golang.org/x/image/colornames"
)

type ScreenSize struct {
	Height int
	Width  int
}

type ScreenConfig struct {
	Size    ScreenSize `json:"size" yaml:"size"`
	Palette []string   `json:"palette" yaml:"palette"`
}

type ImageConfig struct {
	Color string `json:"color" yaml:"color"`
}

func (c *ImageConfig) GetColor() color.Color {
	return colornames.Map[c.Color]
}
