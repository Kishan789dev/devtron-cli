package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func MigrateChartStoreApp(request models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	response := models.DeploymentAppTypeChangeResponse{
		EnvId:                 request.EnvId,
		DesiredDeploymentType: request.DesiredDeploymentType,
	}
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return response, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated, please try with a valid token")
		return response, err
	}
	response, err = client.MigrateChartStoreApp(request)
	if err != nil {
		return response, err
	}
	return response, nil
}
