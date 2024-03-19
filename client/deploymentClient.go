package client

import "github.com/devtron-labs/devtron-cli/devtctl/client/models"

func UpdateDeploymentType(request models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	response := models.Response[models.DeploymentAppTypeChangeResponse]{}

	err := CallPostApi(CHANGE_TYPE, request, &response)

	return response.Result, err

}

func TriggerDeploy(request models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	response := models.Response[models.DeploymentAppTypeChangeResponse]{}

	err := CallPostApi(TRIGGER_DEPLOY, request, &response)

	return response.Result, err
}
