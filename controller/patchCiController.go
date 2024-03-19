package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"log"
	"os"
)

func PatchCiPipelines(manifest cipipeline.CPipelineManifest) {

	err := CheckForValidCriteria(manifest)
	if err != nil {
		fmt.Println("Invalid yaml payload configuration", err)
		return
	}
	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return
	}
	if !isUserAuthenticated {
		fmt.Println("User is not authenticated")
		return
	}

	appIdToPipelineIds := make(map[int][]int)

	var appIdToAppName, pipelineIdToName map[int]string

	for index, payload := range manifest.Spec.Payload {
		appIdToPipelineIds, appIdToAppName, pipelineIdToName, err = GetAppIdToCiPipelineIdsForCriteria(payload.Criteria)
		printImpactedPipelines(index, appIdToPipelineIds, appIdToAppName, pipelineIdToName)
		if err != nil {
			fmt.Println("appId to ciPipelineIds mappings not found", err)
			return
		}
		for appId, pipelineIds := range appIdToPipelineIds {
			patchCiPipelines(appId, pipelineIds, payload.PreCiStage, payload.PostCiStage, payload.Criteria.AppendPreCiSteps, payload.Criteria.AppendPostCiSteps)
		}
	}
}

func CheckForValidCriteria(manifest cipipeline.CPipelineManifest) error {
	if len(manifest.Spec.Payload) == 0 {
		return fmt.Errorf("payload is not valid")
	}
	for _, payload := range manifest.Spec.Payload {
		if (payload.Criteria.AppendPreCiSteps || payload.Criteria.AppendPostCiSteps) && len(payload.Criteria.PipelineIds) != 0 {
			if len(payload.Criteria.AppIds) != 0 ||
				len(payload.Criteria.IncludesAppNames) != 0 ||
				len(payload.Criteria.ExcludesAppNames) != 0 ||
				len(payload.Criteria.ProjectNames) != 0 ||
				len(payload.Criteria.EnvironmentNames) != 0 {
				return fmt.Errorf("payload criteria is not valid")
			}
		}
	}
	return nil
}

func printImpactedPipelines(payloadIndex int, appIdToPipelineIds map[int][]int, appIdToAppName map[int]string, pipelineIdToName map[int]string) {
	fmt.Println("impacted pipelines for payload: ", payloadIndex+1)
	for appId, pipelineIds := range appIdToPipelineIds {
		fmt.Println("- ", "app Id:", appId, " app name: ", appIdToAppName[appId])
		for _, id := range pipelineIds {
			fmt.Println("\t- ", "pipeline Id:", id, " pipeline name: ", pipelineIdToName[id])
		}
	}

	if !viper.GetBool("allYes") {
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
	}
}
func patchCiPipelines(appId int, pipelineIds []int, preCiStage cipipeline.PipelineStage, postCiStage cipipeline.PipelineStage, appendPreCiSteps bool, appendPostCiSteps bool) {
	for _, pipelineId := range pipelineIds {
		fmt.Println("Patching for pipeline ID", pipelineId)
		ciPipeline, err := client.GetCiPipelineForAppIdAndPipelineId(appId, pipelineId)
		if err != nil {
			fmt.Println("error in fetching ci Pipeline with error", err)
		} else {
			patchPipelineRequest := cipipeline.CiPatchRequest{
				CiPipeline:    ciPipeline,
				AppId:         appId,
				Action:        cipipeline.UPDATE_PIPELINE,
				AppWorkflowId: ciPipeline.AppWorkflowId,
			}

			if !appendPreCiSteps {
				if (preCiStage.Steps != nil && len(preCiStage.Steps) != 0) || viper.GetBool("overrideEmptySteps") {
					patchPipelineRequest.CiPipeline.PreBuildStage = preCiStage
				}

			} else {
				for _, step := range preCiStage.Steps {
					patchPipelineRequest.CiPipeline.PreBuildStage.Steps = append(patchPipelineRequest.CiPipeline.PreBuildStage.Steps, step)
				}
			}
			if !appendPostCiSteps {
				if (postCiStage.Steps != nil && len(postCiStage.Steps) != 0) || viper.GetBool("overrideEmptySteps") {
					patchPipelineRequest.CiPipeline.PostBuildStage = postCiStage
				}
			} else {
				for _, step := range postCiStage.Steps {
					patchPipelineRequest.CiPipeline.PostBuildStage.Steps = append(patchPipelineRequest.CiPipeline.PostBuildStage.Steps, step)

				}
			}

			err := client.PatchCiPipeline(patchPipelineRequest)
			if err != nil {
				fmt.Println("error in patching ci Pipeline with error", err)
			}
		}
	}
}
