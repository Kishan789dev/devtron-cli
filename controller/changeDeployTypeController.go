package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func ChangeDeployType(payload models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	response := models.DeploymentAppTypeChangeResponse{
		EnvId:                 payload.EnvId,
		DesiredDeploymentType: payload.DesiredDeploymentType,
		SuccessfulPipelines:   nil,
		FailedPipelines:       nil,
	}
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return response, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return response, err
	}
	response, err = client.UpdateDeploymentType(payload)
	if err != nil {
		return response, err
	}
	return response, nil
}
