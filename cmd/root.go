package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const ImageTypeName = "blank"

var rootCmd = &cobra.Command{
	Use:   ImageTypeName,
	Short: "Generates an image with a single color",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
