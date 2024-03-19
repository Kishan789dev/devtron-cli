package client

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"strconv"
)

func GetCiPipelineConfigForApp(appId int) (cipipeline.CiConfig, error) {

	response := models.Response[cipipeline.CiConfig]{}
	err := CallGetApi(CI_PIPELINES+string(strconv.Itoa(appId)), make(map[string]string), &response)
	return response.Result, err
}

func GetCiPipelineForAppIdAndPipelineId(appId int, pipelineId int) (cipipeline.CiPipeline, error) {
	response := models.Response[cipipeline.CiPipeline]{}
	err := CallGetApi(CI_PIPELINES+string(strconv.Itoa(appId))+"/"+string(strconv.Itoa(pipelineId)), make(map[string]string), &response)
	return response.Result, err
}

func PatchCiPipeline(request cipipeline.CiPatchRequest) error {
	response := models.Response[any]{}
	err := CallPostApi(PATCH_CI_PIPELINES, request, &response)
	return err
}

func GetCiPipelineForPipelineId(pipelineId int) (cipipeline.CiPipelineDetails, error) {
	response := models.Response[cipipeline.CiPipelineDetails]{}
	err := CallGetApi(CI_PIPELINES+"?pipelineId="+string(strconv.Itoa(pipelineId)), make(map[string]string), &response)
	return response.Result, err
}
