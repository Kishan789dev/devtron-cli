/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/cmd/promotionPolicyCmd"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devtctl",
	Short: "devtctl is a CLI which offers operations on Devtron ",
	Long: `devtctl is a CLI which offers operations on Devtron, 
you can use our CLI to perform available operations on your configured devtron app`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringP("server-url", "u", "", "Devtron server URL")
	rootCmd.PersistentFlags().StringP("auth-token", "a", "", "API access token")
	rootCmd.PersistentFlags().StringP("path", "p", "", "path for config yaml/kustomize zip file")
	rootCmd.PersistentFlags().StringP("output-path", "o", "", "path for output")

	viper.BindPFlag("server_url", rootCmd.PersistentFlags().Lookup("server-url"))
	viper.BindPFlag("auth_token", rootCmd.PersistentFlags().Lookup("auth-token"))
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
	viper.BindPFlag("output_path", rootCmd.PersistentFlags().Lookup("output-path"))
	rootCmd.AddCommand(promotionPolicyCmd.CreateCmd)
	rootCmd.AddCommand(promotionPolicyCmd.DeleteCmd)
	rootCmd.AddCommand(promotionPolicyCmd.GetCmd)
	rootCmd.AddCommand(promotionPolicyCmd.UpdateCmd)
	rootCmd.AddCommand(promotionPolicyCmd.ApplyCmd)

	viper.SetConfigFile("./devtctl.env")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
}
