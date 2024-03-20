/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetListOfPromotionPoliciesCmd represents the getListOfPromotionPolicies command
var GetListOfPromotionPoliciesCmd = &cobra.Command{
	Use:   "listOfPromotionPolicies",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		promotionPolicyHandler.HandleListOfPolicies()

	},
}

func init() {
	GetListOfPromotionPoliciesCmd.PersistentFlags().String("search", "", "search")
	GetListOfPromotionPoliciesCmd.PersistentFlags().String("sortBy", "", "sortBy")
	GetListOfPromotionPoliciesCmd.PersistentFlags().String("sortOrder", "", "sortOrder")
	GetListOfPromotionPoliciesCmd.PersistentFlags().Bool("expand", false, "expanding the rows")

	viper.BindPFlag("searchPolicyList", GetListOfPromotionPoliciesCmd.PersistentFlags().Lookup("search"))
	viper.BindPFlag("sortByPolicyList", GetListOfPromotionPoliciesCmd.PersistentFlags().Lookup("sortBy"))
	viper.BindPFlag("sortOrderPolicyList", GetListOfPromotionPoliciesCmd.PersistentFlags().Lookup("sortOrder"))
	viper.BindPFlag("expand", GetListOfPromotionPoliciesCmd.PersistentFlags().Lookup("expand"))

	//cmd.PersistentFlags().StringVar(&expandFlag, "expand", "true", "Expanding the rows")

}
