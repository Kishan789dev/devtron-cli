package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/spf13/viper"
)

func HandleAddKustomizeData() {
	zipPath := viper.GetString("path")
	appId := viper.GetInt(AppId)
	if appId <= 0 {
		fmt.Println("Invalid App Id")
		return
	}
	envId := viper.GetInt(EnvId)
	if envId <= 0 {
		fmt.Println("Invalid Environment Id")
		return
	}
	err := controller.AddKustomizeDataInTheActiveChart(appId, envId, zipPath)
	if err != nil {
		fmt.Println(err)
	}
}
