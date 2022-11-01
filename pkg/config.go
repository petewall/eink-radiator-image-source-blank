package pkg

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"

	"github.com/petewall/eink-radiator-image-source-blank/internal"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . ImageGenerator
type ImageGenerator interface {
	GenerateImage(width, height int) image.Image
}

type Config struct {
	Color string `json:"color" yaml:"color"`
}

func (c *Config) GenerateImage(width, height int) image.Image {
	dst := internal.NewImage(width, height)
	color := &image.Uniform{colornames.Map[c.Color]}
	internal.Draw(dst, dst.Rect, color, image.Point{}, draw.Src)

	return dst
}

func ParseConfig(path string) (*Config, error) {
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read image config file: %w", err)
	}

	var config *Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse image config file: %w", err)
	}

	return config, nil
}
