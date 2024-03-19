package models

type ApplicationType string

const (
	DevtronAppType    ApplicationType = "devtron"
	ChartStoreAppType ApplicationType = "chart-store"
)

type DeploymentAppTypeChangeRequest struct {
	EnvId                 int             `yaml:"envId"`
	DesiredDeploymentType string          `yaml:"desiredDeploymentType"`
	ExcludeApps           []int           `yaml:"excludeApps"`
	IncludeApps           []int           `yaml:"includeApps"`
	AutoTriggerDeployment bool            `yaml:"autoTriggerDeployment"`
	AppType               ApplicationType `yaml:"appType" validate:"oneof=devtron chart-store"`
}

type DeploymentChangeStatus struct {
	PipelineId     int    `json:"pipelineId,omitempty"`
	InstalledAppId int    `json:"installedAppId,omitempty"`
	AppId          int    `json:"appId,omitempty"`
	AppName        string `json:"appName,omitempty"`
	EnvId          int    `json:"envId,omitempty"`
	EnvName        string `json:"envName,omitempty"`
	Error          string `json:"error,omitempty"`
	Status         string `json:"status,omitempty"`
}

type DeploymentAppTypeChangeResponse struct {
	EnvId                 int                       `json:"envId,omitempty"`
	DesiredDeploymentType string                    `json:"desiredDeploymentType,omitempty"`
	SuccessfulPipelines   []*DeploymentChangeStatus `json:"successfulPipelines"`
	FailedPipelines       []*DeploymentChangeStatus `json:"failedPipelines"`
}
