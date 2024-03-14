package controller

import (
	"errors"
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	cd_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/cd-pipeline"
	"golang.org/x/exp/slices"
	"os"
	"strings"
)

func DownloadCdConfigController(cdPayload cd_pipeline.CPipelineManifest) (cd_pipeline.CPipelineManifest, error) {

	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return cd_pipeline.CPipelineManifest{}, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return cd_pipeline.CPipelineManifest{}, err
	}

	cdPipelineManifest, err := getPrePostConfigManifestForCd(cdPayload)
	if err != nil {
		return cd_pipeline.CPipelineManifest{}, err
	}

	return cdPipelineManifest, nil
}
func getPrePostConfigManifestForCd(cdPayload cd_pipeline.CPipelineManifest) (cd_pipeline.CPipelineManifest, error) {
	var cdPipelineManifest cd_pipeline.CPipelineManifest
	if len(cdPayload.Spec.Payload) > 1 {
		return cdPipelineManifest, errors.New(`Invalid yaml config for getting pre-post cd configurations, "payload" label must contain only one item `)
	}
	appIdToCdPipelineInfoMapping, err := GetAppIdToCdPipelineIdsForCriteria(cdPayload.Spec.Payload[0].Criteria)
	if err != nil {
		fmt.Println("There was an error while fetching appId and ciPipelineId Mappings", "err:", err)
	}
	if len(appIdToCdPipelineInfoMapping) == 0 {
		return cdPipelineManifest, errors.New("No criteria provided for selection ")
	}
	cdPipelineManifest = cd_pipeline.CPipelineManifest{
		ApiVersion: "v1",
		Kind:       cdPayload.Kind,
		Metadata:   cdPayload.Metadata,
	}
	result, err := client.GetAppListAutocomplete()
	if err != nil {
		fmt.Errorf("error calling app list autocomplete %s", err)
	}
	for appId, values := range appIdToCdPipelineInfoMapping {
		for _, value := range values {
			cdPipeline, err := client.GetCdPipelineV2(appId, value.EnvId)
			if err != nil {
				fmt.Println("error in fetching cd Pipeline with error", err)
			} else {

				criteria := cd_pipeline.Criteria{
					AppIds: []int{appId},
				}
				for _, env := range result.Environments {
					if env.Id == cdPipeline.Pipelines[0].EnvironmentId {
						criteria.EnvironmentNames = []string{env.Environment}
					}
				}

				cdPayload := cd_pipeline.Payload{
					Criteria: criteria,
				}
				preBuildStage := cdPipeline.Pipelines[0].PreDeployStage
				postBuildStage := cdPipeline.Pipelines[0].PostDeployStage
				if len(preBuildStage.Steps) == 0 && len(postBuildStage.Steps) == 0 {
					continue
				}
				if len(preBuildStage.Steps) > 0 {
					cdPayload.PreCdStage = preBuildStage
				}
				if len(postBuildStage.Steps) > 0 {
					cdPayload.PostCdStage = postBuildStage
				}
				cdPipelineManifest.Spec.Payload = append(cdPipelineManifest.Spec.Payload, cdPayload)
			}
		}
	}
	if len(cdPipelineManifest.Spec.Payload) == 0 {
		fmt.Println("No pipelines found with pre-cd or post-cd configured")
		os.Exit(1)
	}
	return cdPipelineManifest, nil
}

func GetAppIdToCdPipelineIdsForCriteria(criteria cd_pipeline.Criteria) (map[int][]pipelineEnv, error) {
	includeAppNames := criteria.IncludesAppNames
	finalAppToEnvPipelineMapping := make(map[int][]pipelineEnv)
	var err error
	appIds := criteria.AppIds
	var _appIds []int
	if len(appIds) > 0 {
		finalAppToEnvPipelineMapping, err = processCDForAppIdsOrEnvIds(appIds, finalAppToEnvPipelineMapping)
		if err != nil {
			return finalAppToEnvPipelineMapping, err
		}
	}
	if len(includeAppNames) > 0 {
		finalAppToEnvPipelineMapping, err = processCdForIncludeAppName(includeAppNames, finalAppToEnvPipelineMapping)
		if err != nil {
			return finalAppToEnvPipelineMapping, err
		}

	}
	if len(criteria.EnvironmentNames) > 0 || len(criteria.ProjectNames) > 0 {

		envIds := make([]int, 0)
		teamIds := make([]int, 0)
		envNameToEnvId := make(map[string]int)
		teamNameToTeamId := make(map[string]int)
		result, err := client.GetAppListAutocomplete()
		if err != nil {
			return nil, fmt.Errorf("error calling app list autocomplete %s", err)
		}
		for _, environment := range result.Environments {
			envNameToEnvId[environment.Environment] = environment.Id
		}
		for _, team := range result.Teams {
			teamNameToTeamId[team.Name] = team.Id
		}
		if len(criteria.EnvironmentNames) > 0 {
			for _, name := range criteria.EnvironmentNames {
				if envId, ok := envNameToEnvId[name]; ok {
					envIds = append(envIds, envId)
				}
			}
		}
		if len(criteria.ProjectNames) > 0 {
			for _, name := range criteria.ProjectNames {
				if teamId, ok := teamNameToTeamId[name]; ok {
					teamIds = append(teamIds, teamId)
				}
			}
		}

		if len(envIds) == 0 && len(teamIds) == 0 {
			return nil, fmt.Errorf("provided env names or project names did not match")
		}
		apps, err := client.GetAppListForEnvsAndTeams(envIds, teamIds)
		if err != nil {
			return nil, fmt.Errorf("error fetchings apps for environment and project ids %s", err)
		}
		for _, container := range apps.AppContainers {
			if len(finalAppToEnvPipelineMapping) > 0 {
				if _, ok := finalAppToEnvPipelineMapping[container.AppId]; !ok {
					continue
				}
			}
			_appIds = append(_appIds, container.AppId)
		}
		if _appIds == nil {
			return nil, fmt.Errorf("invalid combination of app and environment! given env do/doesn't belong to any of the application")
		}
		finalAppToEnvPipelineMapping, err = processCDForAppIdsOrEnvIds(_appIds, finalAppToEnvPipelineMapping, envIds...)
	}

	return finalAppToEnvPipelineMapping, nil
}

type pipelineEnv struct {
	pipelineId int
	EnvId      int
}

func processCDForAppIdsOrEnvIds(appIds []int, finalAppToEnvPipelineMapping map[int][]pipelineEnv, envIds ...int) (map[int][]pipelineEnv, error) {
	for _, appId := range appIds {
		appIdResetFlagForMapping := false
		cdPipelineDetails, err := client.GetCdPipelines(appId)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println("No app found for the appId:", appId)
			continue
		} else if err != nil {
			return finalAppToEnvPipelineMapping, err
		}
		for _, cdPipelineDetail := range cdPipelineDetails.Pipelines {
			if envIds != nil && !slices.Contains(envIds, cdPipelineDetail.EnvironmentId) {
				continue
			}
			if envIds != nil && !appIdResetFlagForMapping {
				delete(finalAppToEnvPipelineMapping, appId)
				appIdResetFlagForMapping = true
			}
			if value, ok := finalAppToEnvPipelineMapping[appId]; ok {
				isMatchFound := isMatchFoundForPipelineId(value, cdPipelineDetail)
				if isMatchFound {
					continue
				}
				value = append(value, pipelineEnv{
					pipelineId: cdPipelineDetail.Id,
					EnvId:      cdPipelineDetail.EnvironmentId,
				})
				finalAppToEnvPipelineMapping[appId] = value
			} else {
				finalAppToEnvPipelineMapping[appId] = []pipelineEnv{{
					pipelineId: cdPipelineDetail.Id,
					EnvId:      cdPipelineDetail.EnvironmentId,
				}}
			}
		}
	}
	return finalAppToEnvPipelineMapping, nil
}

func processCdForIncludeAppName(includeAppNames []string, finalAppToEnvPipelineMapping map[int][]pipelineEnv) (map[int][]pipelineEnv, error) {
	allApps, err := fetchAllAppDetails()
	if err != nil {
		return finalAppToEnvPipelineMapping, err
	}
	var includedAppIds []int

	for _, app := range allApps {
		if canBeIncluded(includeAppNames, app.AppName) {
			includedAppIds = append(includedAppIds, app.AppId)
		}
	}

	if len(includedAppIds) == 0 {
		fmt.Println("No apps found for includesAppNames regex")
		return finalAppToEnvPipelineMapping, nil
	}
	for _, appId := range includedAppIds {
		cdPipelineDetails, err := client.GetCdPipelines(appId)
		if err != nil {
			return finalAppToEnvPipelineMapping, err
		}
		for i, cdPipelineDetail := range cdPipelineDetails.Pipelines {

			if value, ok := finalAppToEnvPipelineMapping[appId]; ok {
				isMatchFound := isMatchFoundForPipelineId(value, cdPipelineDetail)
				if isMatchFound {
					continue
				}
				value = append(value, pipelineEnv{
					pipelineId: cdPipelineDetail.Id,
					EnvId:      cdPipelineDetails.Pipelines[i].EnvironmentId,
				})
				finalAppToEnvPipelineMapping[appId] = value
			} else {
				finalAppToEnvPipelineMapping[appId] = []pipelineEnv{{
					pipelineId: cdPipelineDetail.Id,
					EnvId:      cdPipelineDetails.Pipelines[i].EnvironmentId,
				}}
			}
		}
	}

	return finalAppToEnvPipelineMapping, nil
}

func isMatchFoundForPipelineId(value []pipelineEnv, cdPipelineDetail cd_pipeline.CDPipelineConfigObject) bool {
	isMatchFound := false
	for _, k := range value {
		if k.pipelineId == cdPipelineDetail.Id {
			isMatchFound = true
			break
		}
	}
	return isMatchFound
}
