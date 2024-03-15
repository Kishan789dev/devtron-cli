package controller

import (
	"errors"
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	ci_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"os"
	"regexp"
	"slices"
	"strings"
)

func DownloadCiConfigController(ciPayload ci_pipeline.CPipelineManifest) (ci_pipeline.CPipelineManifest, error) {

	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return ci_pipeline.CPipelineManifest{}, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return ci_pipeline.CPipelineManifest{}, err
	}

	ciPipelineManifest, err := getPrePostConfigManifest(ciPayload)
	if err != nil {
		return ci_pipeline.CPipelineManifest{}, err
	}

	return ciPipelineManifest, nil
}

func getPrePostConfigManifest(ciPayload ci_pipeline.CPipelineManifest) (ci_pipeline.CPipelineManifest, error) {
	var ciPipelineManifest ci_pipeline.CPipelineManifest
	if len(ciPayload.Spec.Payload) > 1 {
		return ciPipelineManifest, errors.New(`Invalid yaml config for getting pre-post ci configurations, "payload" label must contain only one item `)
	}
	appIdCiPipelineIdsMapping, _, _, err := GetAppIdToCiPipelineIdsForCriteria(ciPayload.Spec.Payload[0].Criteria)
	if err != nil {
		fmt.Println("There was an error while fetching appId and ciPipelineId Mappings", "err:", err)
	}
	if len(appIdCiPipelineIdsMapping) == 0 {
		return ciPipelineManifest, errors.New("No criteria provided for selection ")
	}
	ciPipelineManifest = ci_pipeline.CPipelineManifest{
		ApiVersion: "v1",
		Kind:       ciPayload.Kind,
		Metadata:   ciPayload.Metadata,
	}

	for appId, ciPipelineIds := range appIdCiPipelineIdsMapping {
		for _, ciPipelineId := range ciPipelineIds {
			ciPipeline, err := client.GetCiPipelineForAppIdAndPipelineId(appId, ciPipelineId)
			if err != nil {
				fmt.Println("error in fetching ci Pipeline with error", err)
			} else {
				if ciPipeline.IsExternal == true || ciPipeline.AppType == models.Job {
					continue
				}
				criteria := ci_pipeline.Criteria{
					PipelineIds: []int{ciPipelineId},
				}
				ciPayload := ci_pipeline.Payload{
					Criteria: criteria,
				}
				preBuildStage := ciPipeline.PreBuildStage
				postBuildStage := ciPipeline.PostBuildStage
				if len(preBuildStage.Steps) == 0 && len(postBuildStage.Steps) == 0 {
					continue
				}
				if len(preBuildStage.Steps) > 0 {
					ciPayload.PreCiStage = preBuildStage
				}
				if len(postBuildStage.Steps) > 0 {
					ciPayload.PostCiStage = postBuildStage
				}
				ciPipelineManifest.Spec.Payload = append(ciPipelineManifest.Spec.Payload, ciPayload)
			}
		}
	}
	if len(ciPipelineManifest.Spec.Payload) == 0 {
		fmt.Println("No pipelines found with pre-ci or post-ci configured")
		os.Exit(1)
	}
	return ciPipelineManifest, nil
}

func GetAppIdToCiPipelineIdsForCriteria(criteria ci_pipeline.Criteria) (map[int][]int, map[int]string, map[int]string, error) {
	appIdCiPipelineIdsMappingForAllCriteria := make(map[int][]int)
	mappingForPipelineIdCriteria := make(map[int][]int)
	mappingForAppIdCriteria := make(map[int][]int)
	mappingForAppNameRegexCriteria := make(map[int][]int)

	appIdAppNameMapping := make(map[int]string)
	ciPipelineIdAndNameMapping := make(map[int]string)

	var err error

	ciPipelineIds := criteria.PipelineIds
	appIds := criteria.AppIds
	//includePipelineNames := criteria.IncludesPipelineNames
	//excludePipelineNames := criteria.ExcludesPipelineNames
	includeAppNames := criteria.IncludesAppNames
	excludeAppNames := criteria.ExcludesAppNames

	if len(criteria.EnvironmentNames) > 0 || len(criteria.ProjectNames) > 0 {

		envIds := make([]int, 0)
		teamIds := make([]int, 0)
		envNameToEnvId := make(map[string]int)
		teamNameToTeamId := make(map[string]int)
		result, err := client.GetAppListAutocomplete()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error calling app list autocomplete %s", err)
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
			return nil, nil, nil, fmt.Errorf("provided env names or project names did not match")
		}
		apps, err := client.GetAppListForEnvsAndTeams(envIds, teamIds)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("error fetchings apps for environment and project ids %s", err)
		}
		var _appIds []int
		for _, container := range apps.AppContainers {
			//extract the intersection of appIds , includeAppNames,apps and append the intersection to local variable(_appIds).
			if slices.Contains(appIds, container.AppId) || slices.Contains(includeAppNames, container.AppName) {
				if !slices.Contains(_appIds, container.AppId) {
					_appIds = append(_appIds, container.AppId)
				}
			}
			appIds = append(appIds, container.AppId)
		}
		if appIds != nil || includeAppNames != nil {
			//this is the case when user will provide environmentName and appId or includeAppNames.
			ciPipelineIds, err = getCiPipelineIdForEnv(envIds, _appIds)
			if ciPipelineIds != nil {
				appIds = nil
				includeAppNames = nil
			} else {
				//this the case when user will provide environmentName only.
				ciPipelineIds, err = getCiPipelineIdForEnv(envIds, appIds)
				appIds = nil
				includeAppNames = nil
			}
		}
	}

	if len(ciPipelineIds) > 0 {
		mappingForPipelineIdCriteria, err = processForCiPipelineIds(ciPipelineIds, &appIdAppNameMapping, &ciPipelineIdAndNameMapping)
		if err != nil {
			return appIdCiPipelineIdsMappingForAllCriteria, appIdAppNameMapping, ciPipelineIdAndNameMapping, err
		}
	}
	if len(appIds) > 0 {
		mappingForAppIdCriteria, err = processForAppIds(appIds, &appIdAppNameMapping, &ciPipelineIdAndNameMapping)
		if err != nil {
			return appIdCiPipelineIdsMappingForAllCriteria, appIdAppNameMapping, ciPipelineIdAndNameMapping, err
		}
	}
	if len(includeAppNames) > 0 {
		mappingForAppNameRegexCriteria, err = processForIncludeAppName(includeAppNames, &appIdAppNameMapping, &ciPipelineIdAndNameMapping)
		if err != nil {
			return appIdCiPipelineIdsMappingForAllCriteria, appIdAppNameMapping, ciPipelineIdAndNameMapping, err
		}
	}
	//fmt.Println("")
	clubAppIdAndCiPipelineMapping(&appIdCiPipelineIdsMappingForAllCriteria, mappingForPipelineIdCriteria, mappingForAppIdCriteria, mappingForAppNameRegexCriteria)

	if len(excludeAppNames) > 0 {
		excludeMappingsForExcludeAppNames(&appIdCiPipelineIdsMappingForAllCriteria, excludeAppNames, appIdAppNameMapping)
	}

	if len(ciPipelineIdAndNameMapping) == 0 {
		fmt.Println("No ci-pipelines found for the given criteria")
		os.Exit(1)
	}
	return appIdCiPipelineIdsMappingForAllCriteria, appIdAppNameMapping, ciPipelineIdAndNameMapping, nil
}

func getCiPipelineIdForEnv(envIds []int, appIds []int) ([]int, error) {
	var _ciPipelineIds []int
	var ciPipelineIds []int
	for _, _appId := range appIds {
		cdPipeline, err := client.GetCdPipelines(_appId)
		if err != nil {
			fmt.Println("error in fetching cdPipeline", err)
		}
		for _, pipeline := range cdPipeline.Pipelines {
			if slices.Contains(envIds, pipeline.EnvironmentId) {
				_ciPipelineIds = append(_ciPipelineIds, pipeline.CiPipelineId)
			}
		}
	}
	ciPipelineIds = append(ciPipelineIds, _ciPipelineIds...)

	return ciPipelineIds, nil
}

func clubAppIdAndCiPipelineMapping(clubbedMapping *map[int][]int, mappingForPipelineIdCriteria map[int][]int,
	mappingForAppIdCriteria map[int][]int, mappingForAppNameRegexCriteria map[int][]int) {

	// helper function to add unique elements to the merged array
	addUniqueElements := func(target []int, elements ...int) []int {
		uniqueElements := make(map[int]bool)
		for _, element := range target {
			uniqueElements[element] = true
		}
		for _, element := range elements {
			if !uniqueElements[element] {
				target = append(target, element)
				uniqueElements[element] = true
			}
		}
		return target
	}

	appIdCiPipelineIdMapping := *clubbedMapping
	for appId, ciPipelineIds := range mappingForPipelineIdCriteria {
		if value, ok := appIdCiPipelineIdMapping[appId]; ok {
			newCiPipelineIds := addUniqueElements(value, ciPipelineIds...)
			appIdCiPipelineIdMapping[appId] = newCiPipelineIds
		} else {
			appIdCiPipelineIdMapping[appId] = ciPipelineIds
		}
	}
	for appId, ciPipelineIds := range mappingForAppIdCriteria {
		if value, ok := appIdCiPipelineIdMapping[appId]; ok {
			newCiPipelineIds := addUniqueElements(value, ciPipelineIds...)
			appIdCiPipelineIdMapping[appId] = newCiPipelineIds
		} else {
			appIdCiPipelineIdMapping[appId] = ciPipelineIds
		}
	}
	for appId, ciPipelineIds := range mappingForAppNameRegexCriteria {
		if value, ok := appIdCiPipelineIdMapping[appId]; ok {
			newCiPipelineIds := addUniqueElements(value, ciPipelineIds...)
			appIdCiPipelineIdMapping[appId] = newCiPipelineIds
		} else {
			appIdCiPipelineIdMapping[appId] = ciPipelineIds
		}
	}
	clubbedMapping = &appIdCiPipelineIdMapping
}

func excludeMappingsForExcludeAppNames(appIdCiPipelineIdsMappingForAllCriteria *map[int][]int, excludeAppNames []string, appIdAppNameMapping map[int]string) {
	for appId, appName := range appIdAppNameMapping {
		if canBeExcluded(excludeAppNames, appName) {
			//remove the corresponding appId from appIdCiPipelineIdsMappingForAllCriteria
			delete(*appIdCiPipelineIdsMappingForAllCriteria, appId)
		}
	}
}

func processForAppIds(appIds []int, appIdAppNameMapping *map[int]string, ciPipelineIdAndNameMapping *map[int]string) (map[int][]int, error) {
	appIdCiPipelineIdsMapping := make(map[int][]int)
	appIdAndNameMapping := *appIdAppNameMapping
	ciPipelineIdNameMapping := *ciPipelineIdAndNameMapping

	for _, appId := range appIds {
		ciPipelineDetails, err := client.GetCiPipelineConfigForApp(appId)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println("No app found for the appId:", appId)
			continue
		} else if err != nil {
			return appIdCiPipelineIdsMapping, err
		}
		for _, ciPipelineDetail := range ciPipelineDetails.CiPipelines {
			if ciPipelineDetail.IsExternal == true || ciPipelineDetail.AppType == models.Job {
				continue
			}
			if value, ok := appIdCiPipelineIdsMapping[appId]; ok {
				value = append(value, ciPipelineDetail.Id)
				appIdCiPipelineIdsMapping[appId] = value
			} else {
				appIdCiPipelineIdsMapping[appId] = []int{ciPipelineDetail.Id}
			}
			if _, ok := ciPipelineIdNameMapping[ciPipelineDetail.Id]; !ok {
				ciPipelineIdNameMapping[ciPipelineDetail.Id] = ciPipelineDetail.Name
			}
		}
		if _, ok := appIdAndNameMapping[ciPipelineDetails.AppId]; !ok {
			appIdAndNameMapping[ciPipelineDetails.AppId] = ciPipelineDetails.AppName
		}
	}
	appIdAppNameMapping = &appIdAndNameMapping
	ciPipelineIdAndNameMapping = &ciPipelineIdNameMapping

	return appIdCiPipelineIdsMapping, nil
}

func processForCiPipelineIds(ciPipelineIds []int, appIdAppNameMapping *map[int]string, ciPipelineIdAndNameMapping *map[int]string) (map[int][]int, error) {
	appIdCiPipelineIdsMapping := make(map[int][]int)
	appIdAndNameMapping := *appIdAppNameMapping
	ciPipelineIdNameMapping := *ciPipelineIdAndNameMapping

	for _, ciPipelineId := range ciPipelineIds {
		ciPipelineDetail, err := client.GetCiPipelineForPipelineId(ciPipelineId)
		if err != nil && strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println("No ciPipeline found for the ciPipelineId:", ciPipelineId)
			continue
		} else if err != nil {
			return appIdCiPipelineIdsMapping, err
		}
		if ciPipelineDetail.IsExternal == true || ciPipelineDetail.AppType == models.Job {
			continue
		}
		if value, ok := appIdCiPipelineIdsMapping[ciPipelineDetail.AppId]; ok {
			value = append(value, ciPipelineId)
			appIdCiPipelineIdsMapping[ciPipelineDetail.AppId] = value
		} else {
			appIdCiPipelineIdsMapping[ciPipelineDetail.AppId] = []int{ciPipelineId}
		}
		if _, ok := appIdAndNameMapping[ciPipelineDetail.AppId]; !ok {
			appIdAndNameMapping[ciPipelineDetail.AppId] = ciPipelineDetail.AppName
		}
		if _, ok := ciPipelineIdNameMapping[ciPipelineDetail.Id]; !ok {
			ciPipelineIdNameMapping[ciPipelineDetail.Id] = ciPipelineDetail.Name
		}
	}
	appIdAppNameMapping = &appIdAndNameMapping
	ciPipelineIdAndNameMapping = &ciPipelineIdNameMapping
	return appIdCiPipelineIdsMapping, nil
}

func processForIncludeAppName(includeAppNames []string, appIdAppNameMapping *map[int]string, ciPipelineIdAndNameMapping *map[int]string) (map[int][]int, error) {
	appIdCiPipelineIdsMapping := make(map[int][]int)
	appIdAndNameMapping := *appIdAppNameMapping
	ciPipelineIdNameMapping := *ciPipelineIdAndNameMapping

	allApps, err := fetchAllAppDetails()
	if err != nil {
		return appIdCiPipelineIdsMapping, err
	}
	var includedAppIds []int

	for _, app := range allApps {
		if canBeIncluded(includeAppNames, app.AppName) {
			includedAppIds = append(includedAppIds, app.AppId)
		}
	}

	if len(includedAppIds) == 0 {
		fmt.Println("No apps found for includesAppNames regex")
		return appIdCiPipelineIdsMapping, nil
	}
	for _, appId := range includedAppIds {
		ciPipelineDetails, err := client.GetCiPipelineConfigForApp(appId)
		if err != nil {
			return appIdCiPipelineIdsMapping, err
		}
		for _, ciPipelineDetail := range ciPipelineDetails.CiPipelines {
			if ciPipelineDetail.IsExternal == true || ciPipelineDetail.AppType == models.Job {
				continue
			}
			if value, ok := appIdCiPipelineIdsMapping[appId]; ok {
				value = append(value, ciPipelineDetail.Id)
				appIdCiPipelineIdsMapping[appId] = value
			} else {
				appIdCiPipelineIdsMapping[appId] = []int{ciPipelineDetail.Id}
			}
			if _, ok := ciPipelineIdNameMapping[ciPipelineDetail.Id]; !ok {
				ciPipelineIdNameMapping[ciPipelineDetail.Id] = ciPipelineDetail.Name
			}
		}
		if _, ok := appIdAndNameMapping[ciPipelineDetails.AppId]; !ok {
			appIdAndNameMapping[ciPipelineDetails.AppId] = ciPipelineDetails.AppName
		}

	}

	appIdAppNameMapping = &appIdAndNameMapping
	ciPipelineIdAndNameMapping = &ciPipelineIdNameMapping

	return appIdCiPipelineIdsMapping, nil
}

func processForIncludePipelineNames(includePipelineNames []string, excludePipelineNames []string,
	appIdAppNameMapping *map[int]string, ciPipelineIdAndNameMapping *map[int]string) (map[int][]int, error) {
	appIdCiPipelineIdsMapping := make(map[int][]int)
	appIdAndNameMapping := *appIdAppNameMapping
	ciPipelineIdNameMapping := *ciPipelineIdAndNameMapping

	allApps, err := fetchAllAppDetails()
	if err != nil {
		return appIdCiPipelineIdsMapping, err
	}
	var includedCiPipelines []ci_pipeline.CiPipeline
	for _, appDetail := range allApps {
		ciPipelineDetails, err := client.GetCiPipelineConfigForApp(appDetail.AppId)
		if err != nil {
			return appIdCiPipelineIdsMapping, err
		}
		for _, ciPipelineDetail := range ciPipelineDetails.CiPipelines {
			if ciPipelineDetail.IsExternal == true || ciPipelineDetail.AppType == models.Job {
				continue
			}
			if canBeIncluded(includePipelineNames, ciPipelineDetail.Name) {
				includedCiPipelines = append(includedCiPipelines, ciPipelineDetail)
			}

		}
	}
	if len(includedCiPipelines) == 0 {
		fmt.Println("No ciPipeline found for includesPipelineNames regex")
		return appIdCiPipelineIdsMapping, nil
	}
	for _, includedCiPipelineDetail := range includedCiPipelines {
		if value, ok := appIdCiPipelineIdsMapping[includedCiPipelineDetail.AppId]; ok {
			value = append(value, includedCiPipelineDetail.Id)
			appIdCiPipelineIdsMapping[includedCiPipelineDetail.AppId] = value
		} else {
			appIdCiPipelineIdsMapping[includedCiPipelineDetail.AppId] = []int{includedCiPipelineDetail.Id}
		}
		if _, ok := appIdAndNameMapping[includedCiPipelineDetail.AppId]; !ok {
			appIdAndNameMapping[includedCiPipelineDetail.AppId] = includedCiPipelineDetail.AppName
		}
		if _, ok := ciPipelineIdNameMapping[includedCiPipelineDetail.Id]; !ok {
			ciPipelineIdNameMapping[includedCiPipelineDetail.Id] = includedCiPipelineDetail.Name
		}
	}
	appIdAppNameMapping = &appIdAndNameMapping
	ciPipelineIdAndNameMapping = &ciPipelineIdNameMapping

	return appIdCiPipelineIdsMapping, nil
}

func fetchAllAppDetails() ([]models.AppNameTypeIdContainer, error) {
	appList, err := client.GetAllAppList()
	return appList, err
}

func canBeIncluded(includeRegex []string, toMatch string) bool {
	shouldBeIncluded := true

	for _, includePipelineName := range includeRegex {
		re, err := regexp.Compile(includePipelineName)
		if err != nil {
			fmt.Println("Bad regex in includesPipelineNames label")
			os.Exit(1)
		}
		if re.MatchString(toMatch) {
			shouldBeIncluded = true
			break
		} else {
			shouldBeIncluded = false
		}
	}

	return shouldBeIncluded
}

func canBeExcluded(excludeRegex []string, toMatch string) bool {
	shouldBeExcluded := false
	for _, excludePipelineName := range excludeRegex {
		re, err := regexp.Compile(excludePipelineName)
		if err != nil {
			fmt.Println("Bad regex in excludesPipelineNames label")
			os.Exit(1)
		}
		if re.MatchString(toMatch) {
			shouldBeExcluded = true
			break
		} else {
			shouldBeExcluded = false
		}
	}
	return shouldBeExcluded
}
