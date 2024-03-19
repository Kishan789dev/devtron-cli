/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getListOfPromotionPoliciesCmd represents the getListOfPromotionPolicies command
var getListOfPromotionPoliciesCmd = &cobra.Command{
	Use:   "getListOfPromotionPolicies",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("getListOfPromotionPolicies called")
		handler.HandleListOfPolicies()

	},
}

func init() {
	getListOfPromotionPoliciesCmd.PersistentFlags().String("search", "", "search")
	getListOfPromotionPoliciesCmd.PersistentFlags().String("sortby", "", "sortby")
	getListOfPromotionPoliciesCmd.PersistentFlags().String("sortorder", "", "sortorder")

	viper.BindPFlag("search", getListOfPromotionPoliciesCmd.PersistentFlags().Lookup("search"))
	viper.BindPFlag("sortby", getListOfPromotionPoliciesCmd.PersistentFlags().Lookup("sortby"))
	viper.BindPFlag("sortorder", getListOfPromotionPoliciesCmd.PersistentFlags().Lookup("sortorder"))

	rootCmd.AddCommand(getListOfPromotionPoliciesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getListOfPromotionPoliciesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getListOfPromotionPoliciesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
