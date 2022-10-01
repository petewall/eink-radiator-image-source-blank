package cmd

import (
	"github.com/spf13/cobra"
)

var Version = "dev"
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of this image source",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("%s version: %s\n", ImageTypeName, Version)
	},
}

func init() {
	rootCmd.AddCommand(VersionCmd)
	VersionCmd.SetOut(VersionCmd.OutOrStdout())
}
