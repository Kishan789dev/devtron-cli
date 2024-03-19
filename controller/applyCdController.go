package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/ghodss/yaml"
)

func ApplyCdPipelines(cdPayload models.CDPayload) models.CDPayload {

	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return models.CDPayload{}
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return models.CDPayload{}
	}

	var failedPipelineConfigs []models.CDPipelineConfig
	for _, pipelineConfig := range cdPayload.CDPipelineConfigs {
		fmt.Printf("Creating pipeline for appID: %d with environmentId: %d\n", pipelineConfig.AppId, cdPayload.EnvironmentId)
		namespace, pipelineExists, err := processPipeline(cdPayload.EnvironmentId, pipelineConfig)
		if !pipelineExists && err != nil {
			failedPipelineConfigs = append(failedPipelineConfigs, pipelineConfig)
			fmt.Printf("Pipeline creation failed with message: %s\n", err)
		} else {
			fmt.Printf("Pipeline created successfully!\n")
			fmt.Printf("Applying overrides if provided!\n")
			processOverrides(cdPayload.EnvironmentId, cdPayload.EnvironmentId, pipelineConfig.Overrides, namespace)
		}
	}

	return models.CDPayload{
		SpecVersion:       cdPayload.SpecVersion,
		EnvironmentId:     cdPayload.EnvironmentId,
		CDPipelineConfigs: failedPipelineConfigs,
	}
}

func processPipeline(envId int, pipelineConfig models.CDPipelineConfig) (string, bool, error) {
	var pipelineName string
	if pipelineConfig.PipelineName != "" {
		pipelineName = pipelineConfig.PipelineName
	} else {
		suggestedName, err := client.GetCdSuggestedName(pipelineConfig.AppId)
		if err != nil {
			return "", false, fmt.Errorf("error when fetching suggested name %s", err)
		}
		pipelineName = suggestedName
	}

	environment, err := client.GetEnvDetails(envId)
	if err != nil {
		return "", false, fmt.Errorf("error when fetching environment details from service %s", err)
	}

	pipelineExists := doesPipelineExist(pipelineConfig.AppId, envId)
	if pipelineExists {
		return environment.Namespace, true, nil
	}

	parentCiComponentId, err := getCiComponentIdForWorkflow(pipelineConfig)
	if err != nil {
		return "", false, fmt.Errorf("error when fetching CI component ID for workflow %d with reason %s", pipelineConfig.AppWorkflowId, err)
	}

	strategy, err := getStrategyForApp(pipelineConfig)
	if err != nil {
		return "", false, fmt.Errorf("error when fetching strategy from service %s", err)
	}

	preStageConfig, postStageConfig := getStageConfig(pipelineConfig)

	pipeline := models.Pipeline{
		AppWorkflowId: pipelineConfig.AppWorkflowId,
		EnvironmentId: envId,
		CiPipelineId:  parentCiComponentId,
		TriggerType:   pipelineConfig.TriggerType,
		Name:          pipelineName,
		Namespace:     environment.Namespace,
		Strategy:      []models.Strategy{strategy},
		PreStage: models.StageConfigRequest{
			Config:      preStageConfig,
			TriggerType: pipelineConfig.PreStageConfig.TriggerType,
			Switch:      "config",
		},
		PostStage: models.StageConfigRequest{
			Config:      postStageConfig,
			TriggerType: pipelineConfig.PostStageConfig.TriggerType,
			Switch:      "config",
		},
		PreStageConfigMapSecretNames: models.StageConfigMapSecretNames{
			ConfigMaps: pipelineConfig.PreStageConfig.ConfigMapNames,
			Secrets:    pipelineConfig.PreStageConfig.SecretNames,
		},
		PostStageConfigMapSecretNames: models.StageConfigMapSecretNames{
			ConfigMaps: pipelineConfig.PostStageConfig.ConfigMapNames,
			Secrets:    pipelineConfig.PostStageConfig.SecretNames,
		},
		RunPreStageInEnvironment:  pipelineConfig.PreStageConfig.RunInEnvironment,
		RunPostStageInEnvironment: pipelineConfig.PostStageConfig.RunInEnvironment,
		IsClusterCdActive:         environment.IsClusterCdActive,
		ParentPipelineId:          parentCiComponentId,
		ParentPipelineType:        "CI_PIPELINE",
		DeploymentAppType:         getDeploymentType(pipelineConfig.DeploymentType),
		DeploymentAppCreated:      false,
		DeploymentTemplate:        strategy.DeploymentTemplate,
	}
	cdRequest := models.CDRequest{
		AppId:     pipelineConfig.AppId,
		Pipelines: []models.Pipeline{pipeline},
	}
	err = client.CreateCdPipeline(cdRequest)
	if err != nil {
		return environment.Namespace, false, fmt.Errorf("pipeline creation failed %s", err)
	}
	return environment.Namespace, false, nil
}

func getDeploymentType(deploymentType string) string {
	if deploymentType == "GITOPS" {
		return "argo_cd"
	}
	if deploymentType == "HELM" {
		return "helm"
	}
	return deploymentType
}

func getStageConfig(pipelineConfig models.CDPipelineConfig) (string, string) {
	var preStageConfigJson, postStageConfigJson json.RawMessage
	var preStageConfig, postStageConfig string
	json.Unmarshal([]byte(pipelineConfig.PreStageConfig.Config), &preStageConfigJson)
	json.Unmarshal([]byte(pipelineConfig.PostStageConfig.Config), &postStageConfigJson)

	preStageConfigYaml, _ := yaml.JSONToYAML(preStageConfigJson)
	postStageConfigYaml, _ := yaml.JSONToYAML(postStageConfigJson)
	if preStageConfigJson == nil {
		preStageConfig = ""
	} else {
		preStageConfig = string(preStageConfigYaml)
	}
	if postStageConfigJson == nil {
		postStageConfig = ""
	} else {
		postStageConfig = string(postStageConfigYaml)
	}
	return preStageConfig, postStageConfig
}

func getStrategyForApp(pipelineConfig models.CDPipelineConfig) (models.Strategy, error) {
	pipelineStrategy, err := client.GetStrategies(pipelineConfig.AppId)
	if err != nil {
		return models.Strategy{}, fmt.Errorf("error when fetching strategies from service %s", err)
	}

	for _, strategy := range pipelineStrategy.Strategy {
		if strategy.DeploymentTemplate == pipelineConfig.DeploymentStrategy {
			strategy.Default = true
			return strategy, nil
		}
	}

	return models.Strategy{}, fmt.Errorf("strategy %s don't exist for app %d", pipelineConfig.DeploymentStrategy, pipelineConfig.AppId)
}

func getCiComponentIdForWorkflow(pipelineConfig models.CDPipelineConfig) (int, error) {
	workflows, err := client.GetWorkflows(pipelineConfig.AppId)
	if err != nil {
		return 0, fmt.Errorf("error when fetching workflows from service %s", err)
	}

	for _, workflow := range workflows.Workflows {
		if workflow.Id == pipelineConfig.AppWorkflowId {
			return findCiComponentId(workflow.Tree)
		}
	}
	return 0, fmt.Errorf("workflow Id doesn't exist for appId %d", pipelineConfig.AppId)
}

func findCiComponentId(trees []models.Tree) (int, error) {
	for _, tree := range trees {
		if tree.Type == "CI_PIPELINE" {
			return tree.ComponentId, nil
		}
	}
	return 0, errors.New("couldn't find parent CI pipeline for workflow")
}

func doesPipelineExist(appId int, envId int) bool {
	pipeline, _ := client.GetCdPipeline(appId, envId)
	if len(pipeline.Pipelines) > 0 {
		return true
	}
	return false
}
