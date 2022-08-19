package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const ImageTypeName = "blank"

var rootCmd = &cobra.Command{
	Use:   ImageTypeName,
	Short: "A brief description of your application",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
