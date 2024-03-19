/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
)

// addEnvCmd represents the addEnvCmd command
var addEnvCmd = &cobra.Command{
	Use:   "addEnv",
	Short: "addEnv will add environment to  cluster id",
	Long:  `addEnv will add environment to cluster using cluster id `,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleAddEnv()
	},
}

func init() {
	rootCmd.AddCommand(addEnvCmd)
}
