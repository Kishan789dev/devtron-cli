package models

type CDPayload struct {
	SpecVersion       string             `yaml:"specVersion"` //defaulted to V1
	EnvironmentId     int                `yaml:"environmentId" validate:"required"`
	CDPipelineConfigs []CDPipelineConfig `yaml:"cdPipelineConfigs" json:"cdPipelineConfigs" validate:"required,dive"`
}

type CDPipelineConfig struct {
	Overrides          `yaml:",inline"`
	AppWorkflowId      int         `yaml:"appWorkflowId" validate:"required"`
	TriggerType        string      `yaml:"triggerType" json:"triggerType" validate:"oneof=MANUAL AUTOMATIC"`
	DeploymentStrategy string      `yaml:"deploymentStrategy" json:"deploymentStrategy" validate:"oneof=CANARY RECREATE ROLLING BLUE-GREEN"`
	PreStageConfig     StageConfig `yaml:"preStageConfig"`
	PostStageConfig    StageConfig `yaml:"postStageConfig"`
	DeploymentType     string      `yaml:"deploymentType" json:"deploymentType" validate:"omitempty,oneof=HELM GITOPS "`
	PipelineName       string      `yaml:"pipelineName" json:"pipelineName" validate:"omitempty"`
}
type Overrides struct {
	AppId                          int                 `yaml:"appId"`
	AppName                        string              `yaml:"appName"`
	ChartRefVersion                string              `yaml:"chartRefVersion"`
	DeploymentTemplateOverrideJson string              `yaml:"deploymentTemplateOverrideJson"`
	NewSecrets                     []ConfigMapOrSecret `yaml:"newSecrets"`
	NewConfigs                     []ConfigMapOrSecret `yaml:"newConfigs"`
	ExistingSecrets                []ConfigMapOrSecret `yaml:"existingSecrets"`
	ExistingConfigs                []ConfigMapOrSecret `yaml:"existingConfigs"`
	IsClone                        bool
}

type StageConfig struct {
	Config           string   `yaml:"config"`
	TriggerType      string   `yaml:"triggerType"`
	ConfigMapNames   []string `yaml:"configMapNames"`
	SecretNames      []string `yaml:"secretNames"`
	RunInEnvironment bool     `yaml:"runInEnvironment"`
}

type ConfigMapOrSecret struct {
	IsExternal     bool           `yaml:"isExternal"`
	Name           string         `yaml:"name"`
	KeyValue       []KeyValuePair `yaml:"keyValue"`
	Type           string         `yaml:"type"`
	SubPath        bool           `yaml:"subPath"`
	MountPath      string         `yaml:"mountPath"`
	FilePermission string         `yaml:"filePermission"`
	SubPathKeys    []string       `yaml:"subPathKeys"`
}

type KeyValuePair struct {
	Key   string `yaml:"key" json:"key"`
	Value string `yaml:"value" json:"value"`
}

type CDRequest struct {
	AppId     int        `json:"appId"`
	Pipelines []Pipeline `json:"pipelines"`
}

type Pipeline struct {
	AppWorkflowId                 int                       `json:"appWorkflowId"`
	EnvironmentId                 int                       `json:"environmentId"`
	CiPipelineId                  int                       `json:"ciPipelineId"`
	TriggerType                   string                    `json:"triggerType"`
	Name                          string                    `json:"name"`
	Namespace                     string                    `json:"namespace"`
	Strategy                      []Strategy                `json:"strategies"`
	PreStage                      StageConfigRequest        `json:"preStage"`
	PostStage                     StageConfigRequest        `json:"postStage"`
	PreStageConfigMapSecretNames  StageConfigMapSecretNames `json:"preStageConfigMapSecretNames"`
	PostStageConfigMapSecretNames StageConfigMapSecretNames `json:"postStageConfigMapSecretNames"`
	RunPreStageInEnvironment      bool                      `json:"runPreStageInEnvironment"`
	RunPostStageInEnvironment     bool                      `json:"runPostStageInEnvironment"`
	IsClusterCdActive             bool                      `json:"isClusterCdActive"`
	ParentPipelineId              int                       `json:"parentPipelineId"`
	ParentPipelineType            string                    `json:"parentPipelineType"`
	DeploymentAppType             string                    `json:"deploymentAppType"`
	DeploymentAppCreated          bool                      `json:"deploymentAppCreated"`
	DeploymentTemplate            string                    `json:"deploymentTemplate"`
	AppId                         int                       `json:"appId"`
}

type StageConfigRequest struct {
	Config      string `json:"config"`
	TriggerType string `json:"triggerType"`
	Switch      string `json:"switch"`
}
type StageConfigMapSecretNames struct {
	ConfigMaps []string `json:"configMaps"`
	Secrets    []string `json:"secrets"`
}
