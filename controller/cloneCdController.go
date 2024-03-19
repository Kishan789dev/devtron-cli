package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"regexp"
)

func CloneCdPipelines(cdPayload models.CDClonePayload) models.CDClonePayload {

	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return models.CDClonePayload{}
	}

	if !isUserAuthenticated {
		fmt.Println("User is not authenticated")
		return models.CDClonePayload{}
	}

	appNameToAppId, err := getAppNameToId(cdPayload, err)
	if err != nil {
		fmt.Println(err.Error())
		return models.CDClonePayload{}
	}

	fmt.Println("Impacted App Names")
	//fetching App IDs
	var appIds []int
	if cdPayload.RunForALlApps {
		for appName, appId := range appNameToAppId {
			matched, err := isAppNameMatched(appName, cdPayload.IncludesAppName, cdPayload.ExcludesAppName)
			if err != nil {
				fmt.Printf(err.Error())
				return models.CDClonePayload{}
			}
			if matched {
				appIds = append(appIds, appId)
				fmt.Println(appName)
			}
		}
	} else {
		for _, override := range cdPayload.Overrides {
			if appId, ok := appNameToAppId[override.AppName]; ok {
				appIds = append(appIds, appId)
			}
			fmt.Println(override.AppName)
		}
	}

	prompt := promptui.Select{
		Label: "Do you want to continue? [Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	if result == "No" {
		os.Exit(1)
	}

	//fetching existing CD pipelines in environment
	appIdToPipeline, err := getPipelinesInEnvironment(cdPayload.SourceEnvId)
	if err != nil {
		fmt.Println("Couldn't retrieve CD pipelines for given source environment", err)
		return models.CDClonePayload{}
	}

	//fetching workflow Ids for existing CD pipelines
	appIdToWorkflowId, err := getAppIdToWorkflowId(cdPayload.SourceEnvId)
	if err != nil {
		fmt.Println("Couldn't retrieve App workflows for source environment", err)
		return models.CDClonePayload{}
	}

	appIdToOverride := getAppIdToOverrides(cdPayload, appNameToAppId)

	//cloning pipelines and applying overrides
	for _, appId := range appIds {
		fmt.Printf("Cloning pipeline for appID: %d with new environmentId: %d\n", appId, cdPayload.EnvironmentId)
		namespace, err := clonePipeline(cdPayload.EnvironmentId, appIdToPipeline[appId], appIdToWorkflowId[appId])
		if err != nil {
			fmt.Printf("Pipeline creation failed with message: %s\n", err)
		} else {
			fmt.Printf("Pipeline created successfully!\n")
		}

		fmt.Printf("Creating overides. \n")
		//defaulting to common if not specified
		overrides := cdPayload.CommonOverrides
		overrides.AppId = appId
		overrides.IsClone = true
		//reading specified override for app
		if value, ok := appIdToOverride[appId]; ok {
			overrides = value
		}
		processOverrides(cdPayload.SourceEnvId, cdPayload.EnvironmentId, overrides, namespace)
	}
	return models.CDClonePayload{}
}

func getAppIdToOverrides(cdPayload models.CDClonePayload, appNameToAppId map[string]int) map[int]models.Overrides {
	appIdToOverride := make(map[int]models.Overrides)
	for _, override := range cdPayload.Overrides {

		var chartRefVersion, deploymentTemplateOverrideJson string
		if override.ChartRefVersion == "" {
			chartRefVersion = cdPayload.CommonOverrides.ChartRefVersion
		} else {
			chartRefVersion = override.ChartRefVersion
		}
		if override.DeploymentTemplateOverrideJson == "" {
			deploymentTemplateOverrideJson = cdPayload.CommonOverrides.DeploymentTemplateOverrideJson
		} else {
			deploymentTemplateOverrideJson = override.DeploymentTemplateOverrideJson
		}
		newSecrets := append(override.NewSecrets, cdPayload.CommonOverrides.NewSecrets...)
		newConfigs := append(override.NewConfigs, cdPayload.CommonOverrides.NewConfigs...)
		existingSecrets := append(override.ExistingSecrets, cdPayload.CommonOverrides.ExistingSecrets...)
		existingConfigs := append(override.ExistingConfigs, cdPayload.CommonOverrides.ExistingConfigs...)

		appId := appNameToAppId[override.AppName]
		finalOverride := models.Overrides{
			AppId:                          appId,
			ChartRefVersion:                chartRefVersion,
			DeploymentTemplateOverrideJson: deploymentTemplateOverrideJson,
			NewSecrets:                     newSecrets,
			NewConfigs:                     newConfigs,
			ExistingSecrets:                existingSecrets,
			ExistingConfigs:                existingConfigs,
			IsClone:                        true,
		}
		appIdToOverride[appId] = finalOverride
	}
	return appIdToOverride
}

func getAppNameToId(cdPayload models.CDClonePayload, err error) (map[string]int, error) {
	list, err := client.GetAppList(cdPayload.SourceEnvId, cdPayload.ProjectIds)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve App List %s", err)
	}

	appNameToAppId := make(map[string]int)
	for _, app := range list.AppContainers {
		appNameToAppId[app.AppName] = app.AppId
	}
	return appNameToAppId, nil
}

func clonePipeline(envId int, pipeline models.Pipeline, workflowId int) (string, error) {

	pipelineName, err := client.GetCdSuggestedName(pipeline.AppId)
	if err != nil {
		return "", fmt.Errorf("error when fetching suggested name %s", err)
	}

	environment, err := client.GetEnvDetails(envId)
	if err != nil {
		return "", fmt.Errorf("error when fetching environment details from service %s", err)
	}

	strategy, err := getStrategy(pipeline.AppId, pipeline.DeploymentTemplate)
	if err != nil {
		return "", fmt.Errorf("error when fetching strategy from service %s", err)
	}

	newPipeline := models.Pipeline{
		AppWorkflowId: workflowId,
		EnvironmentId: envId,
		CiPipelineId:  pipeline.CiPipelineId,
		TriggerType:   pipeline.TriggerType,
		Name:          pipelineName,
		Namespace:     environment.Namespace,
		Strategy:      []models.Strategy{strategy},
		PreStage: models.StageConfigRequest{
			Config:      pipeline.PreStage.Config,
			TriggerType: pipeline.PreStage.TriggerType,
			Switch:      "config",
		},
		PostStage: models.StageConfigRequest{
			Config:      pipeline.PostStage.Config,
			TriggerType: pipeline.PostStage.TriggerType,
			Switch:      "config",
		},
		PreStageConfigMapSecretNames: models.StageConfigMapSecretNames{
			ConfigMaps: pipeline.PreStageConfigMapSecretNames.ConfigMaps,
			Secrets:    pipeline.PreStageConfigMapSecretNames.Secrets,
		},
		PostStageConfigMapSecretNames: models.StageConfigMapSecretNames{
			ConfigMaps: pipeline.PostStageConfigMapSecretNames.ConfigMaps,
			Secrets:    pipeline.PostStageConfigMapSecretNames.Secrets,
		},
		RunPreStageInEnvironment:  false,
		RunPostStageInEnvironment: false,
		IsClusterCdActive:         pipeline.IsClusterCdActive,
		ParentPipelineId:          pipeline.ParentPipelineId,
		ParentPipelineType:        pipeline.ParentPipelineType,
		DeploymentAppType:         pipeline.DeploymentAppType,
		DeploymentAppCreated:      pipeline.DeploymentAppCreated,
		DeploymentTemplate:        pipeline.DeploymentTemplate,
		AppId:                     pipeline.AppId,
	}
	cdRequest := models.CDRequest{
		AppId:     pipeline.AppId,
		Pipelines: []models.Pipeline{newPipeline},
	}

	err = client.CreateCdPipeline(cdRequest)
	if err != nil {
		return environment.Namespace, fmt.Errorf("pipeline creation failed %s", err)
	}
	return environment.Namespace, nil
}

func getStrategy(appId int, strategyName string) (models.Strategy, error) {
	pipelineStrategy, err := client.GetStrategies(appId)
	if err != nil {
		return models.Strategy{}, fmt.Errorf("error when fetching strategies from service %s", err)
	}

	for _, strategy := range pipelineStrategy.Strategy {
		if strategy.DeploymentTemplate == strategyName {
			strategy.Default = true
			return strategy, nil
		}
	}

	return models.Strategy{}, fmt.Errorf("strategy %s don't exist for app %d", strategyName, appId)
}

func getPipelinesInEnvironment(envId int) (map[int]models.Pipeline, error) {
	CDPipelines, err := client.GetCdPipelinesInEnvironment(envId)
	if err != nil {
		return nil, err
	}
	appIdToPipeline := make(map[int]models.Pipeline)
	for _, pipeline := range CDPipelines.Pipelines {
		if pipeline.EnvironmentId == envId {
			appIdToPipeline[pipeline.AppId] = pipeline
		}
	}
	return appIdToPipeline, nil
}

func getAppIdToWorkflowId(envId int) (map[int]int, error) {
	appWorkflow, err := client.GetWorkflowsInEnvironment(envId)
	if err != nil {
		return nil, err
	}

	appIdToWorkflowId := make(map[int]int)
	for _, workflow := range appWorkflow.Workflows {
		appIdToWorkflowId[workflow.AppId] = workflow.Id
	}
	return appIdToWorkflowId, nil
}

func isAppNameMatched(appName string, includesRegex string, excludesRegex string) (bool, error) {

	includesMatch := true
	excludesMatch := false
	var err error
	if includesRegex != "" {
		includesMatch, err = regexp.MatchString(includesRegex, appName)
		if err != nil {
			return false, fmt.Errorf("Invalid includes regex for app name: %s\n", err)
		}
	}
	if excludesRegex != "" {
		excludesMatch, err = regexp.MatchString(excludesRegex, appName)
		if err != nil {
			return false, fmt.Errorf("Invalid excludes regex for app name: %s\n", err)
		}
	}

	return includesMatch && !excludesMatch, nil
}
