package ci_pipeline

type PipelineStage struct {
	Id          int                 `json:"id" yaml:"id"`
	Name        string              `json:"name" yaml:"name"`
	Description string              `json:"description" yaml:"description"`
	Type        PipelineStageType   `json:"type" yaml:"type" validate:"oneof=PRE_CI POST_CI PRE_CD POST_CD"`
	Steps       []PipelineStageStep `json:"steps" yaml:"steps"`
}

type PipelineStageStep struct {
	Id                       int                 `json:"id" yaml:"id"`
	Name                     string              `json:"name" yaml:"name"`
	Description              string              `json:"description" yaml:"description"`
	Index                    int                 `json:"index" yaml:"index"`
	StepType                 PipelineStepType    `json:"stepType" validate:"oneof=INLINE REF_PLUGIN" yaml:"stepType"`
	OutputDirectoryPath      []string            `json:"outputDirectoryPath" yaml:"outputDirectoryPath"`
	InlineStepDetail         InlineStepDetail    `json:"inlineStepDetail" yaml:"inlineStepDetail"`
	RefPluginStepDetail      RefPluginStepDetail `json:"pluginRefStepDetail" yaml:"pluginRefStepDetail"`
	TriggerIfParentStageFail bool                `json:"triggerIfParentStageFail" yaml:"triggerIfParentStageFail"`
}

type InlineStepDetail struct {
	ScriptType               ScriptType                `json:"scriptType" validate:"oneof=SHELL DOCKERFILE CONTAINER_IMAGE" yaml:"scriptType"`
	Script                   string                    `json:"script" yaml:"script"`
	StoreScriptAt            string                    `json:"storeScriptAt" yaml:"storeScriptAt"`
	DockerfileExists         bool                      `json:"dockerfileExists" yaml:"dockerfileExists"`
	MountPath                string                    `json:"mountPath" yaml:"mountPath"`
	MountCodeToContainer     bool                      `json:"mountCodeToContainer" yaml:"mountCodeToContainer"`
	MountCodeToContainerPath string                    `json:"mountCodeToContainerPath" yaml:"mountCodeToContainerPath"`
	MountDirectoryFromHost   bool                      `json:"mountDirectoryFromHost" yaml:"mountDirectoryFromHost"`
	ContainerImagePath       string                    `json:"containerImagePath" yaml:"containerImagePath"`
	ImagePullSecretType      ScriptImagePullSecretType `json:"imagePullSecretType" yaml:"imagePullSecretType" validate:"oneof=CONTAINER_REGISTRY SECRET_PATH"`
	ImagePullSecret          string                    `json:"imagePullSecret" yaml:"imagePullSecret"`
	MountPathMap             []MountPathMap            `json:"mountPathMap" yaml:"mountPathMap"`
	CommandArgsMap           []CommandArgsMap          `json:"commandArgsMap" yaml:"commandArgsMap"`
	PortMap                  []PortMap                 `json:"portMap" yaml:"portMap"`
	InputVariables           []StepVariable            `json:"inputVariables" yaml:"inputVariables"`
	OutputVariables          []StepVariable            `json:"outputVariables" yaml:"outputVariables"`
	ConditionDetails         []ConditionDetail         `json:"conditionDetails" yaml:"conditionDetails"`
}

type RefPluginStepDetail struct {
	PluginId         int               `json:"pluginId" yaml:"pluginId"`
	InputVariables   []StepVariable    `json:"inputVariables" yaml:"inputVariables"`
	OutputVariables  []StepVariable    `json:"outputVariables" yaml:"outputVariables"`
	ConditionDetails []ConditionDetail `json:"conditionDetails" yaml:"conditionDetails"`
}
type MountPathMap struct {
	FilePathOnDisk      string `json:"filePathOnDisk" yaml:"filePathOnDisk"`
	FilePathOnContainer string `json:"filePathOnContainer" yaml:"filePathOnContainer"`
}
type CommandArgsMap struct {
	Command string   `json:"command" yaml:"command"`
	Args    []string `json:"args" yaml:"args"`
}

type PortMap struct {
	PortOnLocal     int `json:"portOnLocal" yaml:"portOnLocal" validate:"number,gt=0"`
	PortOnContainer int `json:"portOnContainer" yaml:"portOnContainer" validate:"number,gt=0"`
}

type StepVariable struct {
	Id                        int                                 `json:"id" yaml:"id"`
	Name                      string                              `json:"name" yaml:"name"`
	Format                    PipelineStageStepVariableFormatType `json:"format" validate:"oneof=STRING NUMBER BOOL DATE" yaml:"format"`
	Description               string                              `json:"description" yaml:"description"`
	IsExposed                 bool                                `json:"isExposed" yaml:"isExposed"`
	AllowEmptyValue           bool                                `json:"allowEmptyValue" yaml:"allowEmptyValue"`
	DefaultValue              string                              `json:"defaultValue" yaml:"defaultValue"`
	Value                     string                              `json:"value" yaml:"value"`
	ValueType                 PipelineStageStepVariableValueType  `json:"variableType" validate:"oneof=NEW FROM_PREVIOUS_STEP GLOBAL" yaml:"variableType"`
	PreviousStepIndex         int                                 `json:"refVariableStepIndex" yaml:"refVariableStepIndex"`
	ReferenceVariableName     string                              `json:"refVariableName" yaml:"refVariableName"`
	VariableStepIndexInPlugin int                                 `json:"variableStepIndexInPlugin" yaml:"variableStepIndexInPlugin"`
	ReferenceVariableStage    PipelineStageType                   `json:"refVariableStage" yaml:"refVariableStage"`
}

type ConditionDetail struct {
	Id                  int                            `json:"id" yaml:"id"`
	ConditionOnVariable string                         `json:"conditionOnVariable" yaml:"conditionOnVariable"` //name of variable on which condition is written
	ConditionType       PipelineStageStepConditionType `json:"conditionType" validate:"oneof=SKIP TRIGGER SUCCESS FAIL" yaml:"conditionType"`
	ConditionalOperator string                         `json:"conditionOperator" yaml:"conditionOperator"`
	ConditionalValue    string                         `json:"conditionalValue" yaml:"conditionalValue"`
}
