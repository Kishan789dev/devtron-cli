package handler

import (
	"fmt"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"gopkg.in/go-playground/validator.v9"
)

func HandleCiApply() {

	var ciPipelineManifest cipipeline.CPipelineManifest
	ciPipelineManifest, err := utils.ReadInputFile(ciPipelineManifest, "path")
	if err != nil {
		ciPipelineManifest, err = utils.ReadInputFileJson(ciPipelineManifest, "path")
		if err != nil {
			fmt.Print("Bad input file", err)
			return
		}
	}

	validate := validator.New()
	err = validate.Struct(ciPipelineManifest)
	if err != nil {
		fmt.Print("Invalid yaml configuration: ", err)
		return
	}

	if ciPipelineManifest.ApiVersion != cipipeline.VERSION_V1 {
		fmt.Println("Invalid version provided in manifest, please use v1 ")
		return
	}

	if ciPipelineManifest.Kind != string(cipipeline.CI_PIPELINE_KIND) {
		fmt.Println("Invalid kind provided in manifest, did you mean CI ")
		return
	}

	if ciPipelineManifest.Metadata.Type == cipipeline.PATCH {
		controller.PatchCiPipelines(ciPipelineManifest)
	} else {
		fmt.Println("Invalid Metadata type provided in manifest")
	}
}
