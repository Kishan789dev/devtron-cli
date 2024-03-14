/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createImagePromotionPolicyCmd represents the imagePromotionPolicy command
var CreateImagePromotionPolicyCmd = &cobra.Command{
	Use:   "imagePromotionPolicy",
	Short: "it will create  image promotion policy",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		promotionPolicyHandler.HandleImagePromotionPolicy()
	},
}

func init() {
	CreateImagePromotionPolicyCmd.PersistentFlags().String("name", "", "name of policy")
	CreateImagePromotionPolicyCmd.PersistentFlags().String("description", "", "description of policy")
	CreateImagePromotionPolicyCmd.PersistentFlags().String("passCondition", "", "pass condition ")
	CreateImagePromotionPolicyCmd.PersistentFlags().String("failCondition", "", "fail condition")
	CreateImagePromotionPolicyCmd.PersistentFlags().Int("approverCount", 0, " approver count")
	CreateImagePromotionPolicyCmd.PersistentFlags().Bool("allowRequestFromApprove", false, "allow request from approve")

	CreateImagePromotionPolicyCmd.PersistentFlags().Bool("allowImageBuilderFromApprove", false, "allow image builder from approve")

	CreateImagePromotionPolicyCmd.PersistentFlags().Bool("allowApproverFromDeploy", false, "allow approver from deploy")
	CreateImagePromotionPolicyCmd.PersistentFlags().String("applyPath", "", "path for yaml file")

	viper.BindPFlag("name", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("descriptionPolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("description"))
	viper.BindPFlag("passConditionPolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("passCondition"))
	viper.BindPFlag("failConditionPolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("failCondition"))
	viper.BindPFlag("approverCountPolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("approverCount"))
	viper.BindPFlag("allowRequestFromApprovePolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("allowRequestFromApprove"))
	viper.BindPFlag("allowImageBuilderFromApprovePolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("allowImageBuilderFromApprove"))
	viper.BindPFlag("allowApproverFromDeployPolicy", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("allowApproverFromDeploy"))
	viper.BindPFlag("applyPath", CreateImagePromotionPolicyCmd.PersistentFlags().Lookup("applyPath"))

}
