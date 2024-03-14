/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/cobra"
)

// applyPolicyCmd represents the applyPolicy command
var ApplyPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "A brief description of your command",
	Long:  `Applying  env and policy names`,
	Run: func(cmd *cobra.Command, args []string) {

		promotionPolicyHandler.HandleApplyPolicy()
	},
}
