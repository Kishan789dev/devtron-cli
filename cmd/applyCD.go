/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// applyCDCmd represents the applyCD command
var applyCDCmd = &cobra.Command{
	Use:   "applyCD",
	Short: "apply CD automates creation of CD pipelines.",
	Long: `apply CD automates creation of CD pipelines for you existing CI pipelines of different apps using 
a yaml declaration`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleCdApply()
	},
}
var patchCD = &cobra.Command{
	Use:   "patchCD",
	Short: " patch CD automates patching of CD pipelines.",
	Long:  `patch CD automates patching of CD pipelines through a yaml based configuration. Use command getCD to download the spec`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleCdPatch()
	},
}

func init() {
	rootCmd.AddCommand(applyCDCmd)
	patchCD.PersistentFlags().Bool("allYesCD", false, "to disable interaction and default to yes")
	patchCD.PersistentFlags().Bool("overrideEmptyStepsCD", false, "flag to enable patch with empty steps, default behavior is to retain steps if there is empty patch")
	viper.BindPFlag("allYesCD", patchCD.PersistentFlags().Lookup("allYesCD"))
	viper.BindPFlag("overrideEmptyStepsCD", patchCD.PersistentFlags().Lookup("overrideEmptyStepsCD"))
	rootCmd.AddCommand(patchCD)

}
