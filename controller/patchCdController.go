package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	cd_pipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/cd-pipeline"
	"github.com/spf13/viper"
)

func PatchCdPipelines(manifest cd_pipeline.CPipelineManifest) {

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

	appIdToPipelineIds := make(map[int][]pipelineEnv)
	for _, payload := range manifest.Spec.Payload {
		appIdToPipelineIds, err = GetAppIdToCdPipelineIdsForCriteria(payload.Criteria)
		if err != nil {
			fmt.Println("appId to cdPipelineIds mappings not found", err)
			return
		}
		for appId, values := range appIdToPipelineIds {
			for _, value := range values {
				patchCdPipelines(appId, value.EnvId, payload.PreCdStage, payload.PostCdStage, payload.Criteria.AppendPreCdSteps, payload.Criteria.AppendPostCdSteps, payload.Criteria.RunPreStageInEnv, payload.Criteria.RunPostStageInEnv)

			}
		}
	}
}

func patchCdPipelines(appId int, envId int, preStage cd_pipeline.PipelineStageDto, postStage cd_pipeline.PipelineStageDto, appendPreCdSteps bool, appendPostCdSteps bool, RunPreStageInEnv bool, RunPostStageInEnv bool) {
	cdPipeline, err := client.GetCdPipelineV2(appId, envId)
	if err != nil {
		fmt.Println("error in fetching cd Pipeline with error", err)
	} else {
		patchPipelineRequest := cd_pipeline.CDPatchRequest{
			Pipeline: &cdPipeline.Pipelines[0],
			AppId:    cdPipeline.AppId,
			Action:   cd_pipeline.CD_UPDATE,
			UserId:   cdPipeline.UserId,
		}
		patchPipelineRequest.Pipeline.RunPreStageInEnv = RunPreStageInEnv
		patchPipelineRequest.Pipeline.RunPostStageInEnv = RunPostStageInEnv
		fmt.Printf("patching for appId: %d and envId: %d\n", appId, patchPipelineRequest.Pipeline.EnvironmentId)
		if !appendPreCdSteps {
			if (preStage.Steps != nil && len(preStage.Steps) != 0) || viper.GetBool("overrideEmptyStepsCD") {

				patchPipelineRequest.Pipeline.PreDeployStage = preStage
			}

		} else {
			for _, step := range preStage.Steps {
				patchPipelineRequest.Pipeline.PreDeployStage.Steps = append(patchPipelineRequest.Pipeline.PreDeployStage.Steps, step)

			}
		}
		if !appendPostCdSteps {
			if (postStage.Steps != nil && len(postStage.Steps) != 0) || viper.GetBool("overrideEmptyStepsCD") {
				patchPipelineRequest.Pipeline.PostDeployStage = postStage
			}
		} else {
			for _, step := range postStage.Steps {
				patchPipelineRequest.Pipeline.PostDeployStage.Steps = append(patchPipelineRequest.Pipeline.PostDeployStage.Steps, step)

			}
		}

		err := client.PatchCDPipeline(patchPipelineRequest)
		if err != nil {
			fmt.Println("error in patching cd Pipeline with error", err)
		}
	}
}
