/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package promotionPolicyCmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler/promotionPolicyHandler"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// GetAppAndEnvListCmd represents the getAppAndEnvList command
var GetAppAndEnvListCmd = &cobra.Command{
	Use:   "imagePromotionEnvList",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		promotionPolicyHandler.HandleGetAppAndEnvList()
	},
}

func init() {
	GetAppAndEnvListCmd.PersistentFlags().String("appNames", "", " app names")
	GetAppAndEnvListCmd.PersistentFlags().String("envNames", "", " env names")
	GetAppAndEnvListCmd.PersistentFlags().String("policyNames", "", " policy names")
	GetAppAndEnvListCmd.PersistentFlags().String("sortBy", "", " sort by applicationName or environmentName, defaults to  applicationName")
	GetAppAndEnvListCmd.PersistentFlags().String("sortOrder", "", " sorting order ASC or DESC , defaults to ASC")
	GetAppAndEnvListCmd.PersistentFlags().String("offset", "", "offset on the filtered result defaults to 0 ")
	GetAppAndEnvListCmd.PersistentFlags().String("size", "", " max size of the filtered result set , defaults to 20")

	viper.BindPFlag("appNamesList", GetAppAndEnvListCmd.PersistentFlags().Lookup("appNames"))
	viper.BindPFlag("envNamesList", GetAppAndEnvListCmd.PersistentFlags().Lookup("envNames"))
	viper.BindPFlag("policyNamesList", GetAppAndEnvListCmd.PersistentFlags().Lookup("policyName"))
	viper.BindPFlag("sortBy", GetAppAndEnvListCmd.PersistentFlags().Lookup("sortBy"))
	viper.BindPFlag("sortOrder", GetAppAndEnvListCmd.PersistentFlags().Lookup("sortOrder"))
	viper.BindPFlag("offset", GetAppAndEnvListCmd.PersistentFlags().Lookup("offset"))
	viper.BindPFlag("size", GetAppAndEnvListCmd.PersistentFlags().Lookup("size"))

}
