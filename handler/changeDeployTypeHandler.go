package handler

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/rodaine/table"
	"gopkg.in/go-playground/validator.v9"
)

func HandleChangeDeployType() {
	var payload models.DeploymentAppTypeChangeRequest
	payload, err := utils.ReadInputFile(payload, "path")
	if err != nil {
		return
	}
	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		fmt.Print("Invalid yaml configuration:", err)
		return
	}
	if payload.AppType != models.DevtronAppType {
		fmt.Println("Cmd not supported for the provided app type, please change the app type to 'devtron' ")
		return
	}
	fmt.Println("Change initiated. Please wait.")
	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Start()
	response, err := controller.ChangeDeployType(payload)
	if err != nil {
		fmt.Println("Error while changing deployment type:- ", err.Error())
		return
	}
	s.Stop()
	fmt.Println("Environment ID:- ", response.EnvId)
	fmt.Println("Desired Deployment Type", response.DesiredDeploymentType)

	fmt.Println("Successful Pipelines")

	tblSuccess := table.New("ID", "EnvId", "EnvName", "AppId", "AppName", "Status")
	tblSuccess = utils.FormatTable(tblSuccess)

	for _, pipeline := range response.SuccessfulPipelines {
		tblSuccess.AddRow(pipeline.PipelineId, pipeline.EnvId, pipeline.EnvName, pipeline.AppId, pipeline.AppName, pipeline.Status)
	}

	tblSuccess.Print()

	if len(response.SuccessfulPipelines) == 0 {
		fmt.Println("Nothing to show here")
	}

	fmt.Println("Failed Pipelines")

	tblFailed := table.New("ID", "EnvId", "EnvName", "AppId", "AppName", "Status")
	tblFailed = utils.FormatTable(tblFailed)

	for _, pipeline := range response.FailedPipelines {
		tblFailed.AddRow(pipeline.PipelineId, pipeline.EnvId, pipeline.EnvName, pipeline.AppId, pipeline.AppName, pipeline.Status)
	}

	tblFailed.Print()

	if len(response.FailedPipelines) == 0 {
		fmt.Println("Nothing to show here")
	}

}
