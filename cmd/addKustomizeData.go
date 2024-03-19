package cmd

import (
	"github.com/devtron-labs/devtron-cli/devtctl/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addKustomizeDataCmd = &cobra.Command{
	Use:   "addKustomize",
	Short: "updates the kustomize files in gitOps",
	Long:  `takes a zip file and uploads the content to the gitOps, In the latest chart`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.HandleAddKustomizeData()
	},
}

func init() {
	addKustomizeDataCmd.PersistentFlags().IntP("app-id", "A", 0, "Application Id in the Devtron system")
	addKustomizeDataCmd.PersistentFlags().IntP("env-id", "E", 0, "Environment Id in which the application is installed")
	viper.BindPFlag("app_id", addKustomizeDataCmd.PersistentFlags().Lookup("app-id"))
	viper.BindPFlag("env_id", addKustomizeDataCmd.PersistentFlags().Lookup("env-id"))
	rootCmd.AddCommand(addKustomizeDataCmd)
}
