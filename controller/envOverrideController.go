package controller

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	jsonpatch "github.com/evanphx/json-patch"
)

func processOverrides(sourceEnvId int, targetEnvId int, overrides models.Overrides, namespace string) {

	if overrides.DeploymentTemplateOverrideJson != "" || overrides.IsClone {
		err := patchAndSaveDeploymentTemplate(sourceEnvId, targetEnvId, overrides, namespace)
		if err != nil {
			fmt.Printf("Deployment template override failed with message: %s\n", err)
		}
	}

	configMapId, configNameToData, secretNameToData, existingConfigs, existingSecrets, err := getConfigMapAndSecrets(overrides.AppId, sourceEnvId, targetEnvId)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if overrides.IsClone {
		overrides.ExistingConfigs = existingConfigs
		overrides.ExistingSecrets = existingSecrets
	}

	if len(overrides.NewConfigs) != 0 || len(overrides.ExistingConfigs) != 0 || overrides.IsClone {
		err := saveConfigMap(targetEnvId, overrides, configNameToData, configMapId)
		if err != nil {
			fmt.Printf("Config map override failed with message: %s\n", err)
		}
	}
	if len(overrides.NewSecrets) != 0 || len(overrides.ExistingSecrets) != 0 || overrides.IsClone {
		err := saveConfigSecrets(targetEnvId, overrides, secretNameToData, configMapId)
		if err != nil {
			fmt.Printf("Secrets override failed with message: %s\n", err)
		}
	}
}

func patchAndSaveDeploymentTemplate(sourceEnvId int, targetEnvId int, overrides models.Overrides, namespace string) error {

	template, chartRefId, chartId, err := getDeploymentTemplate(overrides.AppId, sourceEnvId, overrides.ChartRefVersion, 0)
	if err != nil {
		return err
	}
	if sourceEnvId != targetEnvId {
		_, _, chartId, err = getDeploymentTemplate(overrides.AppId, targetEnvId, overrides.ChartRefVersion, chartRefId)
	}
	if err != nil {
		return err
	}
	var finalTemplate json.RawMessage
	if overrides.DeploymentTemplateOverrideJson != "" {
		var patchJson json.RawMessage
		json.Unmarshal([]byte(overrides.DeploymentTemplateOverrideJson), &patchJson)
		if err != nil {
			return err
		}
		finalTemplate, err = patchTemplate(template, patchJson)
	} else {
		finalTemplate = template
	}

	if err != nil {
		return err
	} else {
		if chartId > 0 {
			err = updateTemplate(targetEnvId, finalTemplate, chartRefId, chartId, namespace)
		} else {
			err = saveTemplate(targetEnvId, overrides, finalTemplate, chartRefId)
		}
	}
	return err
}

func saveConfigSecrets(envId int, overrides models.Overrides, secretNameToData map[string]map[string]string, configMapId int) error {

	configDataRequests := createConfigDataRequest(overrides.NewSecrets, nil, b64.StdEncoding.EncodeToString)
	configDataRequests = append(configDataRequests, createConfigDataRequest(overrides.ExistingSecrets, secretNameToData, b64.StdEncoding.EncodeToString)...)

	for _, request := range configDataRequests {

		if request.External {
			request.ExternalType = "KubernetesSecret"
		}
		_, err := client.SaveConfigSecret(models.ConfigRequest{
			Id:            configMapId,
			AppId:         overrides.AppId,
			EnvironmentId: envId,
			ConfigData:    []models.ConfigDataRequest{request},
		})
		if err != nil {
			return fmt.Errorf("error when saving config secrets %s", err)
		}
	}
	return nil
}

func createConfigDataRequest(configs []models.ConfigMapOrSecret, configNameToData map[string]map[string]string, encoding func([]byte) string) []models.ConfigDataRequest {
	var configDataRequests []models.ConfigDataRequest
	for _, config := range configs {
		var configMap map[string]string
		if configNameToData != nil {
			configMap = configNameToData[config.Name]
		} else {
			configMap = make(map[string]string)
		}

		if config.IsExternal {
			for _, key := range config.SubPathKeys {
				configMap[key] = ""
			}
		} else {
			for _, keyValue := range config.KeyValue {
				configMap[keyValue.Key] = encoding([]byte(keyValue.Value))
			}
		}

		configMapJson, _ := json.Marshal(configMap)
		configDataRequests = append(configDataRequests, models.ConfigDataRequest{
			Name:           config.Name,
			Type:           getConfigDataType(config.Type),
			External:       config.IsExternal,
			Data:           configMapJson,
			SubPath:        config.SubPath,
			MountPath:      config.MountPath,
			FilePermission: config.FilePermission,
		})
	}
	return configDataRequests
}

func saveConfigMap(envId int, overrides models.Overrides, configNameToData map[string]map[string]string, configMapId int) error {

	identity := func(input []byte) string { return string(input) }
	configDataRequests := createConfigDataRequest(overrides.NewConfigs, nil, identity)
	configDataRequests = append(configDataRequests, createConfigDataRequest(overrides.ExistingConfigs, configNameToData, identity)...)

	for _, request := range configDataRequests {
		_, err := client.SaveConfigMap(models.ConfigRequest{
			Id:            configMapId,
			AppId:         overrides.AppId,
			EnvironmentId: envId,
			ConfigData:    []models.ConfigDataRequest{request},
		})
		if err != nil {
			return fmt.Errorf("error when saving config maps %s", err)
		}
	}
	return nil
}

func saveTemplate(envId int, overrides models.Overrides, finalTemplate json.RawMessage, chartRefId int) error {
	_, err := client.SaveDeploymentTemplate(models.SaveEnvironmentPropertiesRequest{
		EnvOverrideValues: finalTemplate,
		EnvironmentId:     envId,
		AppMetrics:        false,
		ChartRefId:        chartRefId,
		IsOverride:        true,
		IsBasicViewLocked: false,
		CurrentViewEditor: "ADVANCED",
	}, overrides.AppId, envId)
	if err != nil {
		return fmt.Errorf("error when saving overides for template %s", err)
	}
	return nil
}

func updateTemplate(envId int, finalTemplate json.RawMessage, chartRefId int, chartId int, namespace string) error {
	_, err := client.UpdateDeploymentTemplate(models.UpdateEnvironmentPropertiesRequest{
		Id:                chartId,
		EnvOverrideValues: finalTemplate,
		Status:            1,
		ManualReviewed:    true,
		Active:            true,
		Namespace:         namespace,
		EnvironmentId:     envId,
		ChartRefId:        chartRefId,
		IsOverride:        true,
		IsBasicViewLocked: false,
		CurrentViewEditor: "ADVANCED",
	})
	if err != nil {
		return fmt.Errorf("error when saving overides for template %s", err)
	}
	return nil
}

func patchTemplate(template json.RawMessage, templatePatch json.RawMessage) (json.RawMessage, error) {
	jsonPatch, err := jsonpatch.DecodePatch(templatePatch)
	if err != nil {
		return nil, fmt.Errorf("invalid json patch in override %s", err)
	}

	finalTemplate, err := jsonPatch.Apply(template)
	if err != nil {
		return nil, fmt.Errorf("couldn't patch template with the provided %s", err)
	}
	return finalTemplate, nil
}

func getDeploymentTemplate(appId int, envId int, chartVersion string, chartRefIdClone int) (json.RawMessage, int, int, error) {
	versions, err := client.GetDeploymentVersions(appId, envId)

	if err != nil {
		return nil, 0, 0, fmt.Errorf("error when fetching deployement versions %s", err)
	}

	var chartRefId int
	if chartVersion == "" {
		if chartRefIdClone > 0 {
			chartRefId = chartRefIdClone
		} else {
			chartRefId = versions.LatestEnvChartRef
		}
	} else {
		for _, ref := range versions.ChartRefs {
			if ref.Version == chartVersion {
				chartRefId = ref.Id
			}
		}
	}

	template, err := client.GetDeploymentTemplate(appId, envId, chartRefId)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error when fetching deployement template %s", err)
	}
	var templateJson json.RawMessage
	if template.EnvironmentConfig.IsOverride {
		templateJson = template.EnvironmentConfig.EnvOverrideValues
	} else {
		templateJson = template.GlobalConfig
	}
	return templateJson, chartRefId, template.EnvironmentConfig.Id, nil
}

func getConfigMapAndSecrets(appId int, sourceEnvId int, targetEnvId int) (int, map[string]map[string]string, map[string]map[string]string, []models.ConfigMapOrSecret, []models.ConfigMapOrSecret, error) {

	configMaps, err := client.GetConfigMaps(appId, sourceEnvId)
	if err != nil {
		return 0, nil, nil, nil, nil, fmt.Errorf("error when fetching config maps %s", err)
	}

	var sourceConfigMapId, targetConfigMapId int
	sourceConfigMapId = configMaps.Id
	if sourceEnvId == targetEnvId {
		targetConfigMapId = sourceConfigMapId
	} else {
		configMaps, err := client.GetConfigMaps(appId, targetEnvId)
		targetConfigMapId = configMaps.Id
		if err != nil {
			return 0, nil, nil, nil, nil, fmt.Errorf("error when fetching config maps %s", err)
		}
	}

	var configNameToData = make(map[string]map[string]string)
	var existingConfigs, existingSecrets []models.ConfigMapOrSecret
	for _, configData := range configMaps.ConfigData {
		valueMap := map[string]string{}
		if string(configData.Data) == "null" {
			json.Unmarshal(configData.DefaultData, &valueMap)
		} else {
			json.Unmarshal(configData.Data, &valueMap)
		}
		configNameToData[configData.Name] = valueMap

		if sourceEnvId != targetEnvId {
			existingConfigs = append(existingConfigs, copyExistingConfigData(valueMap, configData))
		}
	}
	secrets, err := client.GetConfigSecrets(appId, sourceEnvId)
	if err != nil {
		return 0, nil, nil, nil, nil, fmt.Errorf("error when fetching config secrets %s", err)
	}
	var secretNameToData = make(map[string]map[string]string)
	for _, secretData := range secrets.ConfigData {
		valueMap := getSecretsForEdit(appId, sourceEnvId, sourceConfigMapId, secretData.Name)
		secretNameToData[secretData.Name] = valueMap
		if sourceEnvId != targetEnvId {
			existingSecrets = append(existingSecrets, copyExistingConfigData(valueMap, secretData))
		}
	}
	return targetConfigMapId, configNameToData, secretNameToData, existingConfigs, existingSecrets, nil
}

func copyExistingConfigData(valueMap map[string]string, configData models.ConfigDataResponse) models.ConfigMapOrSecret {

	var existingConfigs models.ConfigMapOrSecret
	var keyValues []models.KeyValuePair
	var subPathKeys []string
	for key, value := range valueMap {
		if !configData.External {
			keyValues = append(keyValues, models.KeyValuePair{
				Key:   key,
				Value: value,
			})
		} else {
			subPathKeys = append(subPathKeys, key)
		}
	}
	existingConfigs = models.ConfigMapOrSecret{
		IsExternal:     configData.External,
		Name:           configData.Name,
		KeyValue:       keyValues,
		Type:           getConfigDataTypeInternal(configData.Type),
		SubPath:        configData.SubPath,
		MountPath:      configData.MountPath,
		FilePermission: configData.FilePermission,
		SubPathKeys:    subPathKeys,
	}
	return existingConfigs
}
func getSecretsForEdit(appId int, envId int, configMapId int, secretName string) map[string]string {

	config, _ := client.GetConfigSecretsForEdit(appId, envId, configMapId, secretName)
	valueMap := map[string]string{}
	json.Unmarshal(config.ConfigData[0].Data, &valueMap)
	for key, value := range valueMap {
		decodeString, _ := b64.StdEncoding.DecodeString(value)
		valueMap[key] = string(decodeString)
	}
	return valueMap
}

func getConfigDataType(configDataType string) string {
	if configDataType == "ENVIRONMENT" {
		return "environment"
	} else if configDataType == "VOLUME" {
		return "volume"
	}
	return "environment"
}

func getConfigDataTypeInternal(configDataType string) string {
	if configDataType == "environment" {
		return "ENVIRONMENT"
	} else if configDataType == "volume" {
		return "VOLUME"
	}
	return ""
}
