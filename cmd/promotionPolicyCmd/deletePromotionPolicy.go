/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// DeletePromotionPolicyCmd represents the deletePromotionPolicy command
var DeletePromotionPolicyCmd = &cobra.Command{
	Use:   "promotionPolicy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deletePromotionPolicy called")
		promotionPolicyHandler.HandledeletePromotionPolicy()
	},
}

func init() {
	DeletePromotionPolicyCmd.PersistentFlags().String("name", "", "name of policy")

	viper.BindPFlag("policyNameDelete", DeletePromotionPolicyCmd.PersistentFlags().Lookup("name"))

}
