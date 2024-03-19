package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func TriggerDeploy(payload models.DeploymentAppTypeChangeRequest) models.DeploymentAppTypeChangeResponse {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	response := models.DeploymentAppTypeChangeResponse{
		EnvId:                 payload.EnvId,
		DesiredDeploymentType: payload.DesiredDeploymentType,
		SuccessfulPipelines:   nil,
		FailedPipelines:       nil,
	}
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return response
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return response
	}
	response, err = client.TriggerDeploy(payload)
	if err != nil {
		return response
	}
	return response
}
