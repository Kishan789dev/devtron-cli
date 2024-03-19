/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createImagePromotionPolicyCmd represents the imagePromotionPolicy command
var createImagePromotionPolicyCmd = &cobra.Command{
	Use:   "createImagePromotionPolicy",
	Short: "it will create  image promotion policy",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("imagePromotionPolicy called")
		handler.HandleImagePromotionPolicy()
	},
}

func init() {
	createImagePromotionPolicyCmd.PersistentFlags().String("name", "", "name of policy")
	createImagePromotionPolicyCmd.PersistentFlags().String("description", "", "description of policy")
	createImagePromotionPolicyCmd.PersistentFlags().String("pass_condition", "", "pass condition ")
	createImagePromotionPolicyCmd.PersistentFlags().String("fail_condition", "", "fail condition")
	createImagePromotionPolicyCmd.PersistentFlags().String("approver_count", "", " approver_count")
	createImagePromotionPolicyCmd.PersistentFlags().String("allow_image_builder_from_approve", "", "allow_image_builder_from_approve")
	createImagePromotionPolicyCmd.PersistentFlags().String("allow_request_from_approve", "", "allow_request_from_approve")
	createImagePromotionPolicyCmd.PersistentFlags().String("allow_approver_from_deploy", "", "allow_approver_from_deploy")

	viper.BindPFlag("name", createImagePromotionPolicyCmd.PersistentFlags().Lookup("name"))
	viper.BindPFlag("description", createImagePromotionPolicyCmd.PersistentFlags().Lookup("description"))
	viper.BindPFlag("pass_condition", createImagePromotionPolicyCmd.PersistentFlags().Lookup("pass_condition"))
	viper.BindPFlag("fail_condition", createImagePromotionPolicyCmd.PersistentFlags().Lookup("fail_condition"))
	viper.BindPFlag("approver_count", createImagePromotionPolicyCmd.PersistentFlags().Lookup("approver_count"))
	viper.BindPFlag("allow_request_from_approve", createImagePromotionPolicyCmd.PersistentFlags().Lookup("allow_request_from_approve"))
	viper.BindPFlag("allow_image_builder_from_approve", createImagePromotionPolicyCmd.PersistentFlags().Lookup("allow_image_builder_from_approve"))
	viper.BindPFlag("allow_approver_from_deploy", createImagePromotionPolicyCmd.PersistentFlags().Lookup("allow_approver_from_deploy"))

	//rootCmd.AddCommand(getCICmd)
	rootCmd.AddCommand(createImagePromotionPolicyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createImagePromotionPolicyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createImagePromotionPolicyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
