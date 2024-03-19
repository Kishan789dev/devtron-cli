/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCICmd represents the applyCI command
var getCICmd = &cobra.Command{
	Use:   "getCI",
	Short: "Get CI downloads the yaml/json spec for your specified criteria",
	Long:  `Get CI downloads the yaml/json spec for your specified criteria`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleGetCi()
	},
}

func init() {

	getCICmd.PersistentFlags().String("format", "yaml", "json or yaml. Defaults to yaml")
	getCICmd.PersistentFlags().String("appNames", "", "comma separated list of app names")
	getCICmd.PersistentFlags().String("appIds", "", "comma separated list of app Ids")
	getCICmd.PersistentFlags().String("pipelineIds", "", "comma separated list of pipeline Ids")
	getCICmd.PersistentFlags().String("envNames", "", "comma separated list of environment names")
	getCICmd.PersistentFlags().String("projectNames", "", "comma separated list of project names")

	viper.BindPFlag("format", getCICmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("appNames", getCICmd.PersistentFlags().Lookup("appNames"))
	viper.BindPFlag("appIds", getCICmd.PersistentFlags().Lookup("appIds"))
	viper.BindPFlag("pipelineIds", getCICmd.PersistentFlags().Lookup("pipelineIds"))
	viper.BindPFlag("envNames", getCICmd.PersistentFlags().Lookup("envNames"))
	viper.BindPFlag("projectNames", getCICmd.PersistentFlags().Lookup("projectNames"))
	rootCmd.AddCommand(getCICmd)
}
