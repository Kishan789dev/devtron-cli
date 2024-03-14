package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	cd_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/cd-pipeline"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"strconv"
)

func HandleCdApply() {

	var cdJson models.CDPayload
	cdJson, err := utils.ReadInputFile(cdJson, "path")
	if err != nil {
		return
	}

	refactorCDPayload(&cdJson)

	validate := validator.New()
	err = validate.Struct(cdJson)
	if err != nil {
		fmt.Print("Invalid yaml configuration", err)
		return
	}

	failedPayload := controller.ApplyCdPipelines(cdJson)
	marshal, _ := yaml.Marshal(failedPayload)
	if len(failedPayload.CDPipelineConfigs) != 0 {
		fmt.Println("yaml for failed request, please fix and retry")
		fmt.Println(string(marshal))
	}
}
func HandleCdPatch() {
	var cdPipelineManifest cd_pipeline.CPipelineManifest
	cdPipelineManifest, err := utils.ReadInputFile(cdPipelineManifest, "path")
	if err != nil {
		cdPipelineManifest, err = utils.ReadInputFileJson(cdPipelineManifest, "path")
		if err != nil {
			fmt.Print("Bad input file", err)
			return
		}
	}

	validate := validator.New()
	err = validate.Struct(cdPipelineManifest)
	if err != nil {
		fmt.Print("Invalid yaml configuration: ", err)
		return
	}

	if cdPipelineManifest.ApiVersion != cipipeline.VERSION_V1 {
		fmt.Println("Invalid version provided in manifest, please use v1 ")
		return
	}

	if cdPipelineManifest.Kind != string(cd_pipeline.CD_PIPELINE_KIND) {
		fmt.Println("Invalid kind provided in manifest, did you mean CD ")
		return
	}

	if cdPipelineManifest.Metadata.Type == cipipeline.PATCH {
		controller.PatchCdPipelines(cdPipelineManifest)
	} else {
		fmt.Println("Invalid Metadata type provided in manifest")
	}
}

func HandleGetCd() {

	var cdPipelineManifest cd_pipeline.CPipelineManifest
	var err error
	if viper.GetString("path") == "" {
		cdPipelineManifest, err = getManifestForFlagInputForCd()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {

		cdPipelineManifest, err = utils.ReadInputFile(cdPipelineManifest, "path")
		if err != nil {
			cdPipelineManifest, err = utils.ReadInputFileJson(cdPipelineManifest, "path")
			if err != nil {
				fmt.Print("Bad input file", err)
				return
			}
		}

		validate := validator.New()
		err = validate.Struct(cdPipelineManifest)
		if err != nil {
			fmt.Print("Invalid configuration", err)
			return
		}

		if cdPipelineManifest.ApiVersion != cipipeline.VERSION_V1 {
			fmt.Println("Invalid version provided in manifest, please use v1 ")
			return
		}

		if cdPipelineManifest.Kind != string(cd_pipeline.CD_PIPELINE_KIND) {
			fmt.Println("Invalid kind provided in manifest, did you mean CI-Pipeline ")
			return
		}
	}
	if cdPipelineManifest.Metadata.Type == cipipeline.YAML_DOWNLOAD || cdPipelineManifest.Metadata.Type == cipipeline.JSON_DOWNLOAD {
		response, err := controller.DownloadCdConfigController(cdPipelineManifest)
		if err != nil {
			fmt.Println("Could not download pipeline config: ", err)
			return
		}
		writeManifestToFileForCd(response)
	} else {
		fmt.Println("Invalid Metadata type provided in manifest")
	}
}

func writeManifestToFileForCd(manifest cd_pipeline.CPipelineManifest) {
	var err error
	if manifest.Metadata.Type == cipipeline.YAML_DOWNLOAD {
		err = utils.WriteOutputToFileInYaml(manifest)
	} else if manifest.Metadata.Type == cipipeline.JSON_DOWNLOAD {
		err = utils.WriteOutputToFileInJson(manifest)
	}
	if err != nil {
		fmt.Println("Couldn't write to file ", err)
		return
	}
}

func getManifestForFlagInputForCd() (cd_pipeline.CPipelineManifest, error) {

	appIds := make([]int, 0)
	pipelineIds := make([]int, 0)

	appIdsString := utils.SplitAndTrim(viper.GetString("appIdsCD"))
	appNames := utils.SplitAndTrim(viper.GetString("appNamesCD"))
	envNames := utils.SplitAndTrim(viper.GetString("envNamesCD"))
	projectNames := utils.SplitAndTrim(viper.GetString("projectNamesCD"))

	for i, _ := range appIdsString {
		num, err := strconv.Atoi(appIdsString[i])
		if err != nil {
			return cd_pipeline.CPipelineManifest{}, fmt.Errorf("appId provided is not int")
		}
		appIds = append(appIds, num)
	}

	if len(appIds) == 0 && len(appNames) == 0 && len(pipelineIds) == 0 && len(envNames) == 0 && len(projectNames) == 0 {
		return cd_pipeline.CPipelineManifest{}, fmt.Errorf("no criteria provided")
	}

	criteria := cd_pipeline.Criteria{
		AppIds:           appIds,
		IncludesAppNames: appNames,
		EnvironmentNames: envNames,
		ProjectNames:     projectNames,
	}

	var format cipipeline.ManifestType
	if viper.GetString("formatCD") == "json" {
		format = cipipeline.JSON_DOWNLOAD
	} else {
		format = cipipeline.YAML_DOWNLOAD
	}

	return cd_pipeline.CPipelineManifest{
		ApiVersion: cipipeline.VERSION_V1,
		Kind:       string(cd_pipeline.CD_PIPELINE_KIND),
		Metadata:   cd_pipeline.Metadata{Type: format},
		Spec:       cd_pipeline.Spec{Payload: []cd_pipeline.Payload{{Criteria: criteria}}},
	}, nil
}

func refactorCDPayload(cdjson *models.CDPayload) {
	for index, pipeline := range cdjson.CDPipelineConfigs {
		if pipeline.TriggerType == "" {
			pipeline.TriggerType = "AUTOMATIC"
		}
		if pipeline.DeploymentStrategy == "" {
			pipeline.DeploymentStrategy = "ROLLING"
		}
		cdjson.CDPipelineConfigs[index] = pipeline
	}
}
