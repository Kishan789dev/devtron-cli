package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

func HandleCdClone() {

	var cdJson models.CDClonePayload
	cdJson, err := utils.ReadInputFile(cdJson)
	if err != nil {
		return
	}

	validate := validator.New()
	err = validate.Struct(cdJson)
	if err != nil {
		fmt.Print("Invalid yaml configuration:", err)
		return
	}

	failedPayload := controller.CloneCdPipelines(cdJson)
	yaml.Marshal(failedPayload)

	marshal, _ := yaml.Marshal(failedPayload)
	if len(failedPayload.Overrides) != 0 {
		fmt.Println("yaml for failed request, please fix and retry")
		fmt.Println(string(marshal))
	}
}
