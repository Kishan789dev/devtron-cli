package client

import "github.com/devtron-labs/devtron-cli/devtctl/client/models"

func MigrateChartStoreApp(request models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	response := models.Response[models.DeploymentAppTypeChangeResponse]{}

	err := CallPostApi(MIGRATE_CHART_STORE_APP, request, &response)

	return response.Result, err

}

func TriggerChartStoreApp(request models.DeploymentAppTypeChangeRequest) (models.DeploymentAppTypeChangeResponse, error) {
	response := models.Response[models.DeploymentAppTypeChangeResponse]{}

	err := CallPostApi(TRIGGER_CHART_STORE_APP, request, &response)

	return response.Result, err

}
