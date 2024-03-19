package controller

import "github.com/devtron-labs/devtron-cli/devtctl/client"

func AddKustomizeDataInTheActiveChart(appId int, envId int, zipPath string) error {
	return client.AddKustomizeDataInZip(appId, envId, zipPath)
}
