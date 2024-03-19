/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCDCmd represents the getCD command
var getCDCmd = &cobra.Command{
	Use:   "getCD",
	Short: "Get CD downloads the yaml/json spec for your specified criteria",
	Long:  `Get CD downloads the yaml/json spec for your specified criteria`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleGetCd()
	},
}

func init() {
	getCDCmd.PersistentFlags().String("formatCD", "yaml", "json or yaml. Defaults to yaml")
	getCDCmd.PersistentFlags().String("appNamesCD", "", "comma separated list of app names")
	getCDCmd.PersistentFlags().String("appIdsCD", "", "comma separated list of app Ids")
	getCDCmd.PersistentFlags().String("envNamesCD", "", "comma separated list of environment names")
	getCDCmd.PersistentFlags().String("projectNamesCD", "", "comma separated list of project names")

	viper.BindPFlag("formatCD", getCDCmd.PersistentFlags().Lookup("formatCD"))
	viper.BindPFlag("appNamesCD", getCDCmd.PersistentFlags().Lookup("appNamesCD"))
	viper.BindPFlag("appIdsCD", getCDCmd.PersistentFlags().Lookup("appIdsCD"))
	viper.BindPFlag("envNamesCD", getCDCmd.PersistentFlags().Lookup("envNamesCD"))
	viper.BindPFlag("projectNamesCD", getCDCmd.PersistentFlags().Lookup("projectNamesCD"))
	rootCmd.AddCommand(getCDCmd)
}
