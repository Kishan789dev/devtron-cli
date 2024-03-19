package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

// cloneCDCmd represents the cloneCD command
var cloneCDCmd = &cobra.Command{
	Use:   "cloneCD",
	Short: "cloneCD automates cloning of CD pipelines for given environment.",
	Long:  `cloneCD automates cloning of CD pipelines for you existing CD pipelines of different apps using a yaml declaration`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleCdClone()
	},
}

func init() {
	rootCmd.AddCommand(cloneCDCmd)
}
