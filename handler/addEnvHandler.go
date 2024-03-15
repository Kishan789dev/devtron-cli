package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
)

func HandleAddEnv() {

	var envJson models.EnvConfig
	envJson, err := utils.ReadInputFile(envJson, "path")
	if err != nil {
		return
	}

	validate := validator.New()
	err = validate.Struct(envJson)
	if err != nil {
		fmt.Println("invalid yaml configuration", err)
		return
	}

	failedPayload := controller.AddEnv(envJson)
	marshal, _ := yaml.Marshal(failedPayload)
	if len(failedPayload.EnvPayload) != 0 {
		fmt.Println("yaml for failed request, please fix and retry")
		fmt.Println(string(marshal))
	}
}
