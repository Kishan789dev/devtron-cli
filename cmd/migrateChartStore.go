/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

// changeDeployTypeCmd represents the changeDeployType command
var migrateChartStoreCmd = &cobra.Command{
	Use:   "migrateChartStore",
	Short: "migrateChartStore changes the deployment type for chart store app",
	Long:  `migrateChartStore changes the deployment type in a certain env to desired type either git-ops to helm (or vice-versa)`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.MigrateChartStoreApp()
	},
}

func init() {
	rootCmd.AddCommand(migrateChartStoreCmd)

}
