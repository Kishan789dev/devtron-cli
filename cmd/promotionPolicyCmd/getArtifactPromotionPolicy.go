/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetArtifactPromotionPolicyCmd represents the getArtifactPromotionPolicy command
var GetArtifactPromotionPolicyCmd = &cobra.Command{
	Use:   "artifactPromotionPolicy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		promotionPolicyHandler.HandleGetArtifactPromotionPolicy()
	},
}

func init() {
	GetArtifactPromotionPolicyCmd.PersistentFlags().String("name", "", "name of policy")

	viper.BindPFlag("policyNameOfArtifactPromotionPolicy", GetArtifactPromotionPolicyCmd.PersistentFlags().Lookup("name"))

}
