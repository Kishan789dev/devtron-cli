/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

// changeDeployTypeCmd represents the changeDeployType command
var triggerDeployCmd = &cobra.Command{
	Use:   "triggerDeploy",
	Short: "Trigger Deploy will trigger the deployment",
	Long:  `Trigger Deploy will trigger the deployment for the specific env apps having specific deployment type`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleTriggerDeploy()
	},
}

func init() {
	rootCmd.AddCommand(triggerDeployCmd)

}
