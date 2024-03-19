package cd_pipeline

import "encoding/json"

type CDRequest struct {
	AppId     int                      `json:"appId"`
	Pipelines []CDPipelineConfigObject `json:"pipelines"`
	UserId    int32                    `json:"-"`
}

type CDPipelineConfigObject struct {
	Id                         int                `json:"id,omitempty"  validate:"number" `
	EnvironmentId              int                `json:"environmentId,omitempty"  validate:"number,required" `
	EnvironmentName            string             `json:"environmentName,omitempty" `
	Description                string             `json:"description" validate:"max=40"`
	CiPipelineId               int                `json:"ciPipelineId,omitempty" validate:"number"`
	TriggerType                TriggerType        `json:"triggerType,omitempty" validate:"oneof=AUTOMATIC MANUAL"`
	Name                       string             `json:"name,omitempty" validate:"name-component,max=50"` //pipelineName
	CdArgoSetup                bool               `json:"isClusterCdActive"`
	ParentPipelineId           int                `json:"parentPipelineId"`
	ParentPipelineType         string             `json:"parentPipelineType"`
	DeploymentAppType          string             `json:"deploymentAppType"`
	AppName                    string             `json:"appName"`
	DeploymentAppDeleteRequest bool               `json:"deploymentAppDeleteRequest"`
	DeploymentAppCreated       bool               `json:"deploymentAppCreated"`
	AppId                      int                `json:"appId"`
	TeamId                     int                `json:"-"`
	EnvironmentIdentifier      string             `json:"-" `
	IsVirtualEnvironment       bool               `json:"isVirtualEnvironment"`
	HelmPackageName            string             `json:"helmPackageName"`
	ChartName                  string             `json:"chartName"`
	ChartBaseVersion           string             `json:"chartBaseVersion"`
	ContainerRegistryId        int                `json:"containerRegistryId"`
	RepoUrl                    string             `json:"repoUrl"`
	ManifestStorageType        string             `json:"manifestStorageType"`
	PreDeployStage             PipelineStageDto   `json:"preDeployStage,omitempty"`
	PostDeployStage            PipelineStageDto   `json:"postDeployStage,omitempty"`
	ExternalCiPipelineId       int                `json:"externalCiPipelineId,omitempty"`
	CustomTagObject            *CustomTagData     `json:"customTag"`
	CustomTagStage             *PipelineStageType `json:"customTagStage"`
	EnableCustomTag            bool               `json:"enableCustomTag"`
	SwitchFromCiPipelineId     int                `json:"switchFromCiPipelineId"`
	CDPipelineAddType          CDPipelineAddType  `json:"addType"`
	ChildPipelineId            int                `json:"childPipelineId"`
	RunPreStageInEnv           bool               `json:"runPreStageInEnv,omitempty"`
	RunPostStageInEnv          bool               `json:"runPostStageInEnv,omitempty"`
}

type PreStageConfigMapSecretNames struct {
	ConfigMaps []string `json:"configMaps"`
	Secrets    []string `json:"secrets"`
}

type PostStageConfigMapSecretNames struct {
	ConfigMaps []string `json:"configMaps"`
	Secrets    []string `json:"secrets"`
}

type CdStage struct {
	TriggerType TriggerType `json:"triggerType,omitempty"`
	Name        string      `json:"name,omitempty"`
	Status      string      `json:"status,omitempty"`
	Config      string      `json:"config,omitempty"`
}

type Strategy struct {
	DeploymentTemplate DeploymentStrategy `json:"deploymentTemplate,omitempty"`
	Config             json.RawMessage    `json:"config,omitempty" validate:"string"`
	Default            bool               `json:"default"`
}

type PipelineStageDto struct {
	Id          int                     `yaml:"id" json:"id"`
	Name        string                  `json:"name,omitempty" yaml:"name"`
	Description string                  `json:"description,omitempty" yaml:"description"`
	Type        PipelineStageType       `json:"type,omitempty" validate:"omitempty,oneof=PRE_CI POST_CI PRE_CD POST_CD" yaml:"type"`
	Steps       []*PipelineStageStepDto `json:"steps" yaml:"steps"`
	TriggerType TriggerType             `json:"triggerType,omitempty" yaml:"triggerType"`
}

type CustomTagData struct {
	TagPattern string `json:"tagPattern"`
	CounterX   int    `json:"counterX"`
	Enabled    bool   `json:"enabled"`
}

type PipelineStageStepDto struct {
	Id                       int                     `json:"id" yaml:"id"`
	Name                     string                  `json:"name" yaml:"name"`
	Description              string                  `json:"description" yaml:"description"`
	Index                    int                     `json:"index" yaml:"index"`
	StepType                 PipelineStepType        `json:"stepType" validate:"omitempty,oneof=INLINE REF_PLUGIN" yaml:"stepType"`
	OutputDirectoryPath      []string                `json:"outputDirectoryPath" yaml:"outputDirectoryPath"`
	InlineStepDetail         *InlineStepDetailDto    `json:"inlineStepDetail" yaml:"inlineStepDetail"`
	RefPluginStepDetail      *RefPluginStepDetailDto `json:"pluginRefStepDetail" yaml:"refPluginStepDetail"`
	TriggerIfParentStageFail bool                    `json:"triggerIfParentStageFail" yaml:"triggerIfParentStageFail"`
}

type StepVariableDto struct {
	Id                        int                                 `json:"id" yaml:"id"`
	Name                      string                              `json:"name" yaml:"name"`
	Format                    PipelineStageStepVariableFormatType `json:"format" validate:"oneof=STRING NUMBER BOOL DATE" yaml:"format"`
	Description               string                              `json:"description" yaml:"description"`
	IsExposed                 bool                                `json:"isExposed,omitempty" yaml:"isExposed"`
	AllowEmptyValue           bool                                `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue"`
	DefaultValue              string                              `json:"defaultValue,omitempty" yaml:"defaultValue"`
	Value                     string                              `json:"value" yaml:"value"`
	ValueType                 PipelineStageStepVariableValueType  `json:"variableType,omitempty" validate:"oneof=NEW FROM_PREVIOUS_STEP GLOBAL" yaml:"valueType"`
	PreviousStepIndex         int                                 `json:"refVariableStepIndex,omitempty" yaml:"previousStepIndex"`
	ReferenceVariableName     string                              `json:"refVariableName,omitempty" yaml:"referenceVariableName"`
	VariableStepIndexInPlugin int                                 `json:"variableStepIndexInPlugin,omitempty" yaml:"variableStepIndexInPlugin"`
	ReferenceVariableStage    PipelineStageType                   `json:"refVariableStage" yaml:"referenceVariableStage"`
}

type ConditionDetailDto struct {
	Id                  int                            `json:"id"`
	ConditionOnVariable string                         `json:"conditionOnVariable"` //name of variable on which condition is written
	ConditionType       PipelineStageStepConditionType `json:"conditionType" validate:"oneof=SKIP TRIGGER SUCCESS FAIL"`
	ConditionalOperator string                         `json:"conditionOperator"`
	ConditionalValue    string                         `json:"conditionalValue"`
}

type MountPathMap struct {
	FilePathOnDisk      string `json:"filePathOnDisk"`
	FilePathOnContainer string `json:"filePathOnContainer"`
}

type CommandArgsMap struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

type PortMap struct {
	PortOnLocal     int `json:"portOnLocal" validate:"number,gt=0" yaml:"portOnLocal"`
	PortOnContainer int `json:"portOnContainer" validate:"number,gt=0" yaml:"portOnContainer"`
}

type InlineStepDetailDto struct {
	ScriptType               ScriptType                `json:"scriptType" validate:"omitempty,oneof=SHELL DOCKERFILE CONTAINER_IMAGE" yaml:"scriptType"`
	Script                   string                    `json:"script" yaml:"script"`
	StoreScriptAt            string                    `json:"storeScriptAt" yaml:"storeScriptAt"`
	DockerfileExists         bool                      `json:"dockerfileExists,omitempty" yaml:"dockerfileExists"`
	MountPath                string                    `json:"mountPath,omitempty" yaml:"mountPath"`
	MountCodeToContainer     bool                      `json:"mountCodeToContainer,omitempty" yaml:"mountCodeToContainer"`
	MountCodeToContainerPath string                    `json:"mountCodeToContainerPath,omitempty" yaml:"mountCodeToContainerPath"`
	MountDirectoryFromHost   bool                      `json:"mountDirectoryFromHost" yaml:"mountDirectoryFromHost"`
	ContainerImagePath       string                    `json:"containerImagePath,omitempty" yaml:"containerImagePath"`
	ImagePullSecretType      ScriptImagePullSecretType `json:"imagePullSecretType,omitempty" validate:"omitempty,oneof=CONTAINER_REGISTRY SECRET_PATH" yaml:"imagePullSecretType"`
	ImagePullSecret          string                    `json:"imagePullSecret,omitempty" yaml:"imagePullSecret"`
	MountPathMap             []*MountPathMap           `json:"mountPathMap,omitempty" yaml:"mountPathMap"`
	CommandArgsMap           []*CommandArgsMap         `json:"commandArgsMap,omitempty" yaml:"commandArgsMap"`
	PortMap                  []*PortMap                `json:"portMap,omitempty" yaml:"portMap"`
	InputVariables           []*StepVariableDto        `json:"inputVariables" yaml:"inputVariables"`
	OutputVariables          []*StepVariableDto        `json:"outputVariables" yaml:"outputVariables"`
	ConditionDetails         []*ConditionDetailDto     `json:"conditionDetails" yaml:"conditionDetails"`
}

type RefPluginStepDetailDto struct {
	PluginId         int                   `json:"pluginId" yaml:"pluginId"`
	InputVariables   []*StepVariableDto    `json:"inputVariables" yaml:"inputVariables"`
	OutputVariables  []*StepVariableDto    `json:"outputVariables" yaml:"outputVariables"`
	ConditionDetails []*ConditionDetailDto `json:"conditionDetails" yaml:"conditionDetails"`
}
