package models

import "encoding/json"

type ChartRef struct {
	Id      int    `json:"id"`
	Version string `json:"version"`
}

type DeploymentChartRefs struct {
	ChartRefs         []ChartRef `json:"chartRefs"`
	LatestAppChartRef int        `json:"latestAppChartRef"`
	LatestEnvChartRef int        `json:"latestEnvChartRef"`
}

type EnvironmentPropertiesResponse struct {
	EnvironmentConfig EnvironmentProperties `json:"environmentConfig"`
	GlobalConfig      json.RawMessage       `json:"globalConfig"`
	AppMetrics        bool                  `json:"appMetrics"`
	IsOverride        bool                  `sql:"is_override"`
	GlobalChartRefId  int                   `json:"globalChartRefId,omitempty"  validate:"number"`
	ChartRefId        int                   `json:"chartRefId,omitempty"  validate:"number"`
	Namespace         string                `json:"namespace" validate:"name-space-component"`
	Schema            json.RawMessage       `json:"schema"`
	Readme            string                `json:"readme"`
}

type EnvironmentProperties struct {
	Id                int                  `json:"id"`
	EnvOverrideValues json.RawMessage      `json:"envOverrideValues"`
	Status            ChartStatus          `json:"status" validate:"number,required"` //default new, when its ready for deployment CHARTSTATUS_SUCCESS
	ManualReviewed    bool                 `json:"manualReviewed" validate:"required"`
	Active            bool                 `json:"active" validate:"required"`
	Namespace         string               `json:"namespace" validate:"name-space-component,required"`
	EnvironmentId     int                  `json:"environmentId"`
	EnvironmentName   string               `json:"environmentName"`
	Latest            bool                 `json:"latest"`
	UserId            int32                `json:"-"`
	AppMetrics        bool                 `json:"isAppMetricsEnabled"`
	ChartRefId        int                  `json:"chartRefId,omitempty"  validate:"number"`
	IsOverride        bool                 `json:"isOverride"`
	IsBasicViewLocked bool                 `json:"isBasicViewLocked"`
	CurrentViewEditor ChartsViewEditorType `json:"currentViewEditor"` //default "UNDEFINED" in db
}

type SaveEnvironmentPropertiesRequest struct {
	EnvOverrideValues json.RawMessage      `json:"envOverrideValues"`                       //yaml file
	EnvironmentId     int                  `json:"environmentId"`                           //env-id
	AppMetrics        bool                 `json:"isAppMetricsEnabled"`                     //false
	ChartRefId        int                  `json:"chartRefId,omitempty"  validate:"number"` //chartId
	IsOverride        bool                 `json:"isOverride"`                              //true
	IsBasicViewLocked bool                 `json:"isBasicViewLocked"`                       //false
	CurrentViewEditor ChartsViewEditorType `json:"currentViewEditor"`                       // "ADVANCED"

}

type UpdateEnvironmentPropertiesRequest struct {
	Id                int                  `json:"id"` // id of deployment chart
	EnvOverrideValues json.RawMessage      `json:"envOverrideValues"`
	Status            ChartStatus          `json:"status" validate:"number,required"`                  //default new, when its ready for deployment CHARTSTATUS_SUCCESS
	ManualReviewed    bool                 `json:"manualReviewed" validate:"required"`                 //true
	Active            bool                 `json:"active" validate:"required"`                         //true
	Namespace         string               `json:"namespace" validate:"name-space-component,required"` //"kube-system"
	EnvironmentId     int                  `json:"environmentId"`                                      // env-id
	ChartRefId        int                  `json:"chartRefId,omitempty"  validate:"number"`            //chart id
	IsOverride        bool                 `json:"isOverride"`                                         //true
	IsBasicViewLocked bool                 `json:"isBasicViewLocked"`                                  //false
	CurrentViewEditor ChartsViewEditorType `json:"currentViewEditor"`                                  // "ADVANCED"
}

type ChartStatus int

type ChartsViewEditorType string

type ConfigResponse struct {
	Id            int                  `json:"id"`
	AppId         int                  `json:"appId"`
	EnvironmentId int                  `json:"environmentId,omitempty"`
	ConfigData    []ConfigDataResponse `json:"configData"`
}

type ConfigDataResponse struct {
	Name                 string           `json:"name"`
	Type                 string           `json:"type"`
	External             bool             `json:"external"`
	Data                 json.RawMessage  `json:"data"`
	DefaultData          json.RawMessage  `json:"defaultData"`
	Global               bool             `json:"global"`
	ExternalSecretType   string           `json:"externalType"`
	ESOSecretData        ESOSecretData    `json:"esoSecretData"`
	DefaultESOSecretData ESOSecretData    `json:"defaultESOSecretData,omitempty"`
	ExternalSecret       []ExternalSecret `json:"secretData"`
	RoleARN              string           `json:"roleARN"`
	SubPath              bool             `json:"subPath"`
	FilePermission       string           `json:"filePermission"`
	MountPath            string           `json:"mountPath"`
}

type ConfigRequest struct {
	Id            int                 `json:"id"`
	AppId         int                 `json:"appId"`
	EnvironmentId int                 `json:"environmentId,omitempty"`
	ConfigData    []ConfigDataRequest `json:"configData"`
}

type ConfigDataRequest struct {
	Name           string          `json:"name"`
	Type           string          `json:"type"`
	External       bool            `json:"external"`
	Data           json.RawMessage `json:"data"`
	SubPath        bool            `json:"subPath"`
	MountPath      string          `json:"mountPath"`
	FilePermission string          `json:"filePermission"`
	ExternalType   string          `json:"externalType"`
}

type ESOSecretData struct {
	SecretStore     json.RawMessage `json:"secretStore,omitempty"`
	SecretStoreRef  json.RawMessage `json:"secretStoreRef,omitempty"`
	EsoData         []ESOData       `json:"esoData,omitempty"`
	RefreshInterval string          `json:"refreshInterval,omitempty"`
}
type ESOData struct {
	SecretKey string `json:"secretKey"`
	Key       string `json:"key"`
	Property  string `json:"property,omitempty"`
}

type ExternalSecret struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Property string `json:"property,omitempty"`
	IsBinary bool   `json:"isBinary"`
}
