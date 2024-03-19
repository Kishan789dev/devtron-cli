package client

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	cd_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/cd-pipeline"
	"strconv"
)

func GetEnvDetails(envId int) (models.Environment, error) {

	response := models.Response[models.Environment]{}
	query := map[string]string{"id": strconv.Itoa(envId)}
	err := CallGetApi(Environment_FOR_ENV_ID, query, &response)
	return response.Result, err
}

func GetCdSuggestedName(appId int) (string, error) {

	response := models.Response[string]{}
	err := CallGetApi(CD_SUGGEST_NAME_API+string(strconv.Itoa(appId)), make(map[string]string), &response)
	return response.Result, err
}

func GetWorkflows(appId int) (models.AppWorkflow, error) {

	response := models.Response[models.AppWorkflow]{}
	err := CallGetApi(WORKFLOWS_FOR_APP+string(strconv.Itoa(appId)), make(map[string]string), &response)
	return response.Result, err
}

func GetWorkflowsInEnvironment(envId int) (models.AppWorkflow, error) {

	response := models.Response[models.AppWorkflow]{}
	err := CallGetApi(WORKFLOWS_FOR_ENV+strconv.Itoa(envId)+"/app-wf", make(map[string]string), &response)
	return response.Result, err
}

func GetStrategies(appId int) (models.PipelineStrategy, error) {

	response := models.Response[models.PipelineStrategy]{}
	err := CallGetApi(STRATEGIES_FOR_APP+string(strconv.Itoa(appId)), make(map[string]string), &response)
	return response.Result, err
}

func CreateCdPipeline(request models.CDRequest) error {
	response := models.Response[any]{}
	err := CallPostApi(CREATE_CD_PIPELINE, request, &response)
	return err
}

func GetCdPipeline(appId int, envId int) (models.CDRequest, error) {
	response := models.Response[models.CDRequest]{}
	err := CallGetApi(CD_PIPELINE+strconv.Itoa(appId)+"/env/"+strconv.Itoa(envId), make(map[string]string), &response)
	return response.Result, err
}
func GetCdPipelines(appId int) (cd_pipeline.CDRequest, error) {
	response := models.Response[cd_pipeline.CDRequest]{}
	err := CallGetApi(CD_PIPELINE+strconv.Itoa(appId), make(map[string]string), &response)
	return response.Result, err
}

func GetCdPipelineV2(appId int, envId int) (cd_pipeline.CDRequest, error) {
	response := models.Response[cd_pipeline.CDRequest]{}
	err := CallGetApi(CD_PIPELINE+strconv.Itoa(appId)+"/env/"+strconv.Itoa(envId), make(map[string]string), &response)
	return response.Result, err
}

func GetCdPipelinesInEnvironment(envId int) (models.CDRequest, error) {
	response := models.Response[models.CDRequest]{}
	err := CallGetApi(CD_PIPELINES_ENV+strconv.Itoa(envId)+"/cd-pipeline", make(map[string]string), &response)
	return response.Result, err
}

func GetAppList(envId int, teamIds []int) (models.AppList, error) {
	request := models.AppListRequest{
		Environments: []int{envId},
		Teams:        teamIds,
		Size:         1000,
	}
	response := models.Response[models.AppList]{}
	err := CallPostApi(GET_APPS, request, &response)
	return response.Result, err
}

func GetAppListForEnvsAndTeams(envIds []int, teamIds []int) (models.AppList, error) {
	request := models.AppListRequest{
		Environments: envIds,
		Teams:        teamIds,
		Size:         5000,
	}
	response := models.Response[models.AppList]{}
	err := CallPostApi(GET_APPS, request, &response)
	return response.Result, err
}

func GetAllAppList() ([]models.AppNameTypeIdContainer, error) {
	response := models.Response[[]models.AppNameTypeIdContainer]{}
	err := CallGetApi(ALL_APP_LIST, make(map[string]string), &response)
	return response.Result, err

}
func GetAppListAutocomplete() (models.AppAutocomplete, error) {
	response := models.Response[models.AppAutocomplete]{}
	err := CallGetApi(APP_LIST_AUTOCOMPLETE, make(map[string]string), &response)
	return response.Result, err
}
