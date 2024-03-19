package handler

import (
	"fmt"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
)

func HandleGetCi() {

	var ciPipelineManifest cipipeline.CPipelineManifest
	var err error
	if viper.GetString("path") == "" {
		ciPipelineManifest, err = getManifestForFlagInput()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {

		ciPipelineManifest, err = utils.ReadInputFile(ciPipelineManifest)
		if err != nil {
			ciPipelineManifest, err = utils.ReadInputFileJson(ciPipelineManifest)
			if err != nil {
				fmt.Print("Bad input file", err)
				return
			}
		}

		validate := validator.New()
		err = validate.Struct(ciPipelineManifest)
		if err != nil {
			fmt.Print("Invalid configuration", err)
			return
		}

		if ciPipelineManifest.ApiVersion != cipipeline.VERSION_V1 {
			fmt.Println("Invalid version provided in manifest, please use v1 ")
			return
		}

		if ciPipelineManifest.Kind != string(cipipeline.CI_PIPELINE_KIND) {
			fmt.Println("Invalid kind provided in manifest, did you mean CI-Pipeline ")
			return
		}
	}
	if ciPipelineManifest.Metadata.Type == cipipeline.YAML_DOWNLOAD || ciPipelineManifest.Metadata.Type == cipipeline.JSON_DOWNLOAD {
		response, err := controller.DownloadCiConfigController(ciPipelineManifest)
		if err != nil {
			fmt.Println("Could not download pipeline config: ", err)
			return
		}
		writeManifestToFile(response)
	} else {
		fmt.Println("Invalid Metadata type provided in manifest")
	}
}

func writeManifestToFile(manifest cipipeline.CPipelineManifest) {
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

func splitAndTrim(input string) []string {

	if !strings.Contains(input, ",") {
		return strings.Fields(input)
	}

	splitStrings := strings.Split(input, ",")
	finalSplitStrings := make([]string, 0)

	for _, str := range splitStrings {
		value := strings.TrimSpace(str)
		if value != "" {
			finalSplitStrings = append(finalSplitStrings, value)
		}
	}
	return finalSplitStrings
}

func getManifestForFlagInput() (cipipeline.CPipelineManifest, error) {

	appIds := make([]int, 0)
	pipelineIds := make([]int, 0)

	//appIdsString := strings.Fields(viper.GetString("appIds"))
	//appNames := strings.Fields(viper.GetString("appNames"))
	//envNames := strings.Fields(viper.GetString("envNames"))
	//projectNames := strings.Fields(viper.GetString("projectNames"))
	//pipelineIdsString := strings.Fields(viper.GetString("pipelineIds"))

	appIdsString := splitAndTrim(viper.GetString("appIds"))
	appNames := splitAndTrim(viper.GetString("appNames"))
	envNames := splitAndTrim(viper.GetString("envNames"))
	projectNames := splitAndTrim(viper.GetString("projectNames"))
	pipelineIdsString := splitAndTrim(viper.GetString("pipelineIds"))

	for i, _ := range appIdsString {
		num, err := strconv.Atoi(appIdsString[i])
		if err != nil {
			return cipipeline.CPipelineManifest{}, fmt.Errorf("appId provided is not int")
		}
		appIds = append(appIds, num)
	}

	for i, _ := range pipelineIdsString {
		num, err := strconv.Atoi(pipelineIdsString[i])
		if err != nil {
			return cipipeline.CPipelineManifest{}, fmt.Errorf("pipelineId provided is not int")
		}
		pipelineIds = append(pipelineIds, num)
	}

	if len(appIds) == 0 && len(appNames) == 0 && len(pipelineIds) == 0 && len(envNames) == 0 && len(projectNames) == 0 {
		return cipipeline.CPipelineManifest{}, fmt.Errorf("no criteria provided")
	}

	criteria := cipipeline.Criteria{
		PipelineIds:      pipelineIds,
		AppIds:           appIds,
		IncludesAppNames: appNames,
		EnvironmentNames: envNames,
		ProjectNames:     projectNames,
	}

	var format cipipeline.ManifestType
	if viper.GetString("format") == "json" {
		format = cipipeline.JSON_DOWNLOAD
	} else {
		format = cipipeline.YAML_DOWNLOAD
	}

	return cipipeline.CPipelineManifest{
		ApiVersion: cipipeline.VERSION_V1,
		Kind:       string(cipipeline.CI_PIPELINE_KIND),
		Metadata:   cipipeline.Metadata{Type: format},
		Spec:       cipipeline.Spec{Payload: []cipipeline.Payload{{Criteria: criteria}}},
	}, nil
}
