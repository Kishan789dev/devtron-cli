/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCICmd represents the applyCI command
var applyCICmd = &cobra.Command{
	Use:   "applyCI",
	Short: "apply CI automates patching of CI pipelines",
	Long:  `apply CI automates patching of CI pipelines through a yaml based configuration. Use command getCI to download the spec`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleCiApply()
	},
}

func init() {
	applyCICmd.PersistentFlags().Bool("allYes", false, "to disable interaction and default to yes")
	applyCICmd.PersistentFlags().Bool("overrideEmptySteps", false, "flag to enable patch with empty steps, default behavior is to retain steps if there is empty patch")

	viper.BindPFlag("allYes", applyCICmd.PersistentFlags().Lookup("allYes"))
	viper.BindPFlag("overrideEmptySteps", applyCICmd.PersistentFlags().Lookup("overrideEmptySteps"))

	rootCmd.AddCommand(applyCICmd)
}
