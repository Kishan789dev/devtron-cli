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

func HandleTriggerDeploy() {
	var payload models.DeploymentAppTypeChangeRequest
	payload, err := utils.ReadInputFile(payload)
	if err != nil {
		return
	}
	validate := validator.New()
	err = validate.Struct(payload)
	if err != nil {
		fmt.Print("Invalid yaml configuration:", err)
		return
	}
	fmt.Println(" Deployment Triggered. Please wait.")
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Start()
	response := controller.TriggerDeploy(payload)
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
