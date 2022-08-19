package cmd

import (
	"encoding/json"
	"github.com/petewall/eink-radiator-image-source-blank/m/v2/lib"
	"github.com/spf13/cobra"
)

var blankConfigCmd = &cobra.Command{
	Use:   "blank-config",
	Short: "Print a blank config for the " + ImageTypeName + " image type",
	Run: func(cmd *cobra.Command, args []string) {
		encoded, _ := json.Marshal(lib.ImageConfig{Color: ""})
		cmd.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(blankConfigCmd)
	blankConfigCmd.SetOut(blankConfigCmd.OutOrStdout())
}
