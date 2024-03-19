package ci_pipeline

import (
	"encoding/json"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

type CiConfig struct {
	AppId       int          `json:"appId,omitempty" validate:"required,number"`
	AppName     string       `json:"appName,omitempty"`
	CiPipelines []CiPipeline `json:"ciPipelines,omitempty" validate:"dive"` //a pipeline will be built for each ciMaterial
}

type CiPipeline struct {
	AppWorkflowId              int               `json:"appWorkflowId,omitempty"`
	IsManual                   bool              `json:"isManual"`
	DockerArgs                 map[string]string `json:"dockerArgs"`
	IsExternal                 bool              `json:"isExternal"`
	AppType                    models.AppType    `json:"appType,omitempty"`
	AppId                      int               `json:"appId"`
	AppName                    string            `json:"appName,omitempty"`
	CiMaterial                 json.RawMessage   `json:"ciMaterial,omitempty"`
	Name                       string            `json:"name,omitempty" validate:"name-component,max=100"` //name suffix of corresponding pipeline. required, unique, validation corresponding to gocd pipelineName will be applicable
	Id                         int               `json:"id,omitempty" `
	Active                     bool              `json:"active,omitempty"` //pipeline is active or not
	LinkedCount                int               `json:"linkedCount"`
	ScanEnabled                bool              `json:"scanEnabled,notnull"`
	PreBuildStage              PipelineStage     `json:"preBuildStage,omitempty"`
	PostBuildStage             PipelineStage     `json:"postBuildStage,omitempty"`
	IsDockerConfigOverridden   bool              `json:"isDockerConfigOverridden"`
	DockerConfigOverride       json.RawMessage   `json:"dockerConfigOverride,omitempty"`
	IsOffendingMandatoryPlugin bool              `json:"isOffendingMandatoryPlugin,omitempty"`
}

type CiPipelineDetails struct {
	IsManual                 bool              `json:"isManual"`
	DockerArgs               map[string]string `json:"dockerArgs"`
	IsExternal               bool              `json:"isExternal"`
	ParentCiPipeline         int               `json:"parentCiPipeline"`
	ParentAppId              int               `json:"parentAppId"`
	AppId                    int               `json:"appId"`
	AppName                  string            `json:"appName,omitempty"`
	AppType                  models.AppType    `json:"appType,omitempty"`
	Name                     string            `json:"name,omitempty" validate:"name-component,max=100"` //name suffix of corresponding pipeline. required, unique, validation corresponding to gocd pipelineName will be applicable
	Id                       int               `json:"id,omitempty" `
	Version                  string            `json:"version,omitempty"` //matchIf token version in gocd . used for update request
	Active                   bool              `json:"active,omitempty"`  //pipeline is active or not
	Deleted                  bool              `json:"deleted,omitempty"`
	LinkedCount              int               `json:"linkedCount"`
	ScanEnabled              bool              `json:"scanEnabled,notnull"`
	AppWorkflowId            int               `json:"appWorkflowId,omitempty"`
	TargetPlatform           string            `json:"targetPlatform,omitempty"`
	IsDockerConfigOverridden bool              `json:"isDockerConfigOverridden"`
	EnvironmentId            int               `json:"environmentId,omitempty"`
	LastTriggeredEnvId       int               `json:"lastTriggeredEnvId"`
}
