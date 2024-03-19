package client

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	cd_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/cd-pipeline"
	"strconv"
)

func GetDeploymentVersions(appId int, envId int) (models.DeploymentChartRefs, error) {

	response := models.Response[models.DeploymentChartRefs]{}

	err := CallGetApi(GET_ALL_DEPLOYMENT_TEMPLATE+string(strconv.Itoa(appId))+"/"+string(strconv.Itoa(envId)), make(map[string]string), &response)

	return response.Result, err

}

func GetDeploymentTemplate(appId int, envId int, chartId int) (models.EnvironmentPropertiesResponse, error) {
	response := models.Response[models.EnvironmentPropertiesResponse]{}

	err := CallGetApi(GET_A_DEPLOYMENT_TEMPLATE+"/"+string(strconv.Itoa(appId))+"/"+string(strconv.Itoa(envId))+"/"+string(strconv.Itoa(chartId)), make(map[string]string), &response)

	return response.Result, err
}

func GetConfigMaps(appId int, envId int) (models.ConfigResponse, error) {

	response := models.Response[models.ConfigResponse]{}

	err := CallGetApi(GET_CONFIG_MAP+"/"+string(strconv.Itoa(appId))+"/"+string(strconv.Itoa(envId)), make(map[string]string), &response)

	return response.Result, err
}

func GetConfigSecrets(appId int, envId int) (models.ConfigResponse, error) {

	response := models.Response[models.ConfigResponse]{}

	err := CallGetApi(GET_CONFIG_SECRET+"/"+string(strconv.Itoa(appId))+"/"+strconv.Itoa(envId), make(map[string]string), &response)

	return response.Result, err
}

func GetConfigSecretsForEdit(appId int, envId int, configID int, secretName string) (models.ConfigResponse, error) {

	response := models.Response[models.ConfigResponse]{}

	query := map[string]string{"name": secretName}
	err := CallGetApi(GET_CONFIG_SECRET_EDIT+string(strconv.Itoa(appId))+"/"+strconv.Itoa(envId)+"/"+strconv.Itoa(configID), query, &response)

	return response.Result, err
}

func SaveDeploymentTemplate(request models.SaveEnvironmentPropertiesRequest, appId int, envId int) (any, error) {
	response := models.Response[any]{}

	err := CallPostApi(GET_A_DEPLOYMENT_TEMPLATE+"/"+string(strconv.Itoa(appId))+"/"+string(strconv.Itoa(envId)), request, &response)

	return response.Result, err

}

func UpdateDeploymentTemplate(request models.UpdateEnvironmentPropertiesRequest) (any, error) {
	response := models.Response[any]{}

	err := CallPutApi(GET_A_DEPLOYMENT_TEMPLATE, request, &response)

	return response.Result, err

}

func SaveConfigMap(request models.ConfigRequest) (any, error) {
	response := models.Response[any]{}

	err := CallPostApi(GET_CONFIG_MAP, request, &response)

	return response.Result, err
}

func SaveConfigSecret(request models.ConfigRequest) (any, error) {
	response := models.Response[any]{}

	err := CallPostApi(GET_CONFIG_SECRET, request, &response)

	return response.Result, err
}

func PatchCDPipeline(request cd_pipeline.CDPatchRequest) error {
	response := models.Response[any]{}
	err := CallPostApi(PATCH_CD_PIPELINES, request, &response)
	return err
}
