package internal

import (
	"fmt"
	"image/color"
	"os"

	"golang.org/x/image/colornames"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Color string `json:"color" yaml:"color"`
}

func (c *Config) GetColor() color.Color {
	return colornames.Map[c.Color]
}

func (c *Config) GenerateImage(width, height int) ImageContext {
	imageContext := NewImageContext(width, height)
	imageContext.SetColor(c.GetColor())
	imageContext.DrawRectangle(0, 0, float64(width), float64(height))
	imageContext.Fill()

	return imageContext
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
