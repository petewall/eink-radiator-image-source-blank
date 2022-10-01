package cmd

import (
	"github.com/fogleman/gg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
)

var config *internal.Config

func parseConfig(cmd *cobra.Command, args []string) error {
	var err error
	config, err = internal.ParseConfig(viper.GetString("config"))

	return err
}

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates a " + ImageTypeName + " image",
	PreRunE: parseConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		imageContext := gg.NewContext(viper.GetInt("width"), viper.GetInt("height"))
		imageContext.SetColor(config.GetColor())
		imageContext.DrawRectangle(0, 0, float64(viper.GetInt("width")), float64(viper.GetInt("height")))
		imageContext.Fill()

		if viper.GetBool("to-stdout") {
			return imageContext.EncodePNG(cmd.OutOrStdout())
		}

		return imageContext.SavePNG(viper.GetString("output"))
	},
}

const (
	DefaultImageHeight = 480
	DefaultImageWidth  = 640
)

func init() {
	rootCmd.AddCommand(GenerateCmd)
	GenerateCmd.Flags().StringP("config", "c", "", "the path to the image config file")
	GenerateCmd.Flags().Int("height", DefaultImageHeight, "the height of the image")
	GenerateCmd.Flags().Int("width", DefaultImageWidth, "the width of the image")

	GenerateCmd.Flags().StringP("output", "o", ImageTypeName+".png", "path to write the file")
	GenerateCmd.Flags().Bool("to-stdout", false, "print the image to stdout")
	_ = viper.BindPFlags(GenerateCmd.Flags())
}
