/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

// changeDeployTypeCmd represents the changeDeployType command
var changeDeployTypeCmd = &cobra.Command{
	Use:   "changeDeployType",
	Short: "change deploy type changes the deployment type.",
	Long:  `Change deploy type changes the deployment type in a certain env to desired type either git-ops to helm (or vice-versa) `,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleChangeDeployType()
	},
}

func init() {
	rootCmd.AddCommand(changeDeployTypeCmd)

}
