/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

var triggerChartStoreCmd = &cobra.Command{
	Use:   "triggerChartStore",
	Short: "triggerChartStore triggers the chart store app",
	Long:  `triggerChartStore triggers the chart store app after it has been successfully migrated.`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.TriggerChartStoreApp()
	},
}

func init() {
	rootCmd.AddCommand(triggerChartStoreCmd)

}
