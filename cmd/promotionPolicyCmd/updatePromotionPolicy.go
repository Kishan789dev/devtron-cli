/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// UpdatePromotionPolicyCmd represents the updatePromotionPolicy command
var UpdatePromotionPolicyCmd = &cobra.Command{
	Use:   "imagePromotionPolicy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		promotionPolicyHandler.HandleUpdatePolicy()
	},
}

func init() {
	UpdatePromotionPolicyCmd.PersistentFlags().String("policyName", "", "oldname of policy")
	UpdatePromotionPolicyCmd.PersistentFlags().String("name", "", "newname of policy")
	UpdatePromotionPolicyCmd.PersistentFlags().String("description", "", "description of policy")
	UpdatePromotionPolicyCmd.PersistentFlags().String("passCondition", "", "pass condition ")
	UpdatePromotionPolicyCmd.PersistentFlags().String("failCondition", "", "fail condition")
	UpdatePromotionPolicyCmd.PersistentFlags().String("approverCount", "", " approver_count")
	UpdatePromotionPolicyCmd.PersistentFlags().String("allowImageBuilderFromApprove", "", "allow_image_builder_from_approve")
	UpdatePromotionPolicyCmd.PersistentFlags().String("allowRequestFromApprove", "", "allow_request_from_approve")
	UpdatePromotionPolicyCmd.PersistentFlags().String("allowApproverFromDeploy", "", "allow_approver_from_deploy")

	viper.BindPFlag("policyName", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("policyName"))
	viper.BindPFlag("updatedPolicyName", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("descriptionUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("description"))
	viper.BindPFlag("passConditionUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("passCondition"))
	viper.BindPFlag("failConditionUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("failCondition"))
	viper.BindPFlag("approverCountUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("approverCount"))
	viper.BindPFlag("allowRequestFromApproveUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("allowRequestFromApprove"))
	viper.BindPFlag("allowImageBuilderFromApproveUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("allowImageBuilderFromApprove"))
	viper.BindPFlag("allowApproverFromDeployUpdate", UpdatePromotionPolicyCmd.PersistentFlags().Lookup("allowApproverFromDeploy"))

}
