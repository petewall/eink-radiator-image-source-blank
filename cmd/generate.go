package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/petewall/eink-radiator-image-source-blank/v2/internal"
)

var ImageGenerator internal.ImageGenerator

func parseConfig(cmd *cobra.Command, args []string) error {
	var err error
	ImageGenerator, err = internal.ParseConfig(viper.GetString("config"))
	return err
}

var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generates a " + ImageTypeName + " image",
	PreRunE: parseConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		imageContext := ImageGenerator.GenerateImage(viper.GetInt("width"), viper.GetInt("height"))

		if viper.GetBool("to-stdout") {
			return imageContext.EncodePNG(cmd.OutOrStdout())
		}

		return imageContext.SavePNG(viper.GetString("output"))
	},
}

const (
	DefaultOutputFilename = ImageTypeName + ".png"
)

func init() {
	rootCmd.AddCommand(GenerateCmd)
	GenerateCmd.Flags().StringP("config", "c", "", "the path to the image config file")
	_ = GenerateCmd.MarkFlagRequired("config")
	GenerateCmd.Flags().Int("height", 0, "the height of the image")
	_ = GenerateCmd.MarkFlagRequired("height")
	GenerateCmd.Flags().Int("width", 0, "the width of the image")
	_ = GenerateCmd.MarkFlagRequired("width")

	GenerateCmd.Flags().StringP("output", "o", DefaultOutputFilename, "path to write the file")
	GenerateCmd.Flags().Bool("to-stdout", false, "print the image to stdout")
	GenerateCmd.MarkFlagsMutuallyExclusive("output", "to-stdout")
	GenerateCmd.SetOut(os.Stdout)
	_ = viper.BindPFlags(GenerateCmd.Flags())
}
