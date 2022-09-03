package cmd

import (
	"image/png"
	"os"

	"github.com/fogleman/gg"
	"github.com/petewall/eink-radiator-image-source-blank/v2/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	imageConfig  *lib.ImageConfig
	screenConfig *lib.ScreenConfig
)

func ParseConfigs(cmd *cobra.Command, args []string) error {
	configData, err := os.ReadFile(viper.GetString("config"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configData, &imageConfig)
	if err != nil {
		return err
	}

	configData, err = os.ReadFile(viper.GetString("screen"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configData, &screenConfig)
	if err != nil {
		return err
	}

	return nil
}

var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates a " + ImageTypeName + " image",
	PreRunE: ParseConfigs,
	RunE: func(cmd *cobra.Command, args []string) error {
		imageContext := gg.NewContext(screenConfig.Size.Width, screenConfig.Size.Height)
		imageContext.SetColor(imageConfig.GetColor())
		imageContext.DrawRectangle(0, 0, float64(screenConfig.Size.Width), float64(screenConfig.Size.Height))
		imageContext.Fill()

		if viper.GetBool("to-stdout") {
			return png.Encode(cmd.OutOrStdout(), imageContext.Image())
		}

		return imageContext.SavePNG(viper.GetString("output"))
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("config", "c", "", "the path to the image config file")
	generateCmd.Flags().StringP("screen", "s", "", "the path to the screen config file")

	generateCmd.Flags().StringP("output", "o", ImageTypeName+".png", "path to write the file")
	generateCmd.Flags().Bool("to-stdout", false, "print the image to stdout")
	_ = viper.BindPFlags(generateCmd.Flags())
}
