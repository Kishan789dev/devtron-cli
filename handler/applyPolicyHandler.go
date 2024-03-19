package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/applyPolicy"
	cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

func HandleApplyPolicy() {
	var applyPolicyManifest applyPolicy.ApplyPolicyManifest

	if viper.GetString("path") == "" {
		fmt.Println("provide some path  of the yaml file")
		return

	} else {

		applyPolicyManifest, err := utils.ReadInputFile(applyPolicyManifest)
		if err != nil {
			applyPolicyManifest, err = utils.ReadInputFileJson(applyPolicyManifest)
			if err != nil {
				fmt.Print("Bad input file", err)
				return
			}
		}

		validate := validator.New()
		err = validate.Struct(applyPolicyManifest)
		if err != nil {
			fmt.Print("Invalid configuration", err)
			return
		}
		if applyPolicyManifest.ApiVersion != cipipeline.VERSION_V1 {
			fmt.Println("Invalid version provided in manifest, please use v1 ")
			return
		}

		if applyPolicyManifest.Kind != string(cipipeline.ARTIFACT_PROMOTION_POLICY) {
			fmt.Println("Invalid kind provided in manifest, did you mean CI-Pipeline ")
			return
		}
		fmt.Println(applyPolicyManifest)
		//response ,err:=controller.ApplyPolicyController()

	}
}
