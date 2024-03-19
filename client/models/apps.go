package models

type AppList struct {
	AppContainers []App `json:"appContainers"`
}

type App struct {
	AppId   int    `json:"appId"`
	AppName string `json:"appName"`
}

type AppListRequest struct {
	Environments []int `json:"environments"`
	Teams        []int `json:"teams"`
	Size         int   `json:"size"`
}

type AppNameTypeIdContainer struct {
	AppName string `json:"appName"`
	Type    string `json:"type"`
	AppId   int    `json:"appId"`
}

type AppType int

const (
	CustomApp     AppType = 0 // cicd app
	ChartStoreApp AppType = 1 // helm app
	Job           AppType = 2 // jobs
)

type AppAutocomplete struct {
	Teams        []TeamRequest
	Environments []EnvironmentBean
}

type TeamRequest struct {
	Id   int    `json:"id,omitempty" validate:"number"`
	Name string `json:"name,omitempty" validate:"required"`
}
type EnvironmentBean struct {
	Id                    int    `json:"id,omitempty" validate:"number"`
	Environment           string `json:"environment_name,omitempty" validate:"required,max=50"`
	ClusterId             int    `json:"cluster_id,omitempty" validate:"number,required"`
	ClusterName           string `json:"cluster_name,omitempty"`
	Namespace             string `json:"namespace,omitempty" validate:"name-space-component,max=50"`
	EnvironmentIdentifier string `json:"environmentIdentifier"`
}
