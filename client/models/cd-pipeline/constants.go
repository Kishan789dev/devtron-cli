package cd_pipeline

type TriggerType string //HOW pipeline should be triggered
type PipelineStageType string
type PipelineStepType string
type PipelineStageStepVariableType string
type PipelineStageStepVariableValueType string
type PipelineStageStepConditionType string
type PipelineStageStepVariableFormatType string
type ScriptType string
type ScriptImagePullSecretType string
type DeploymentStrategy string
type CDPipelineAddType string
type CdPatchAction int

const (
	TRIGGER_TYPE_AUTOMATIC TriggerType = "AUTOMATIC"
	TRIGGER_TYPE_MANUAL    TriggerType = "MANUAL"
)
const (
	DEPLOYMENT_STRATEGY_BLUE_GREEN DeploymentStrategy = "BLUE-GREEN"
	DEPLOYMENT_STRATEGY_ROLLING    DeploymentStrategy = "ROLLING"
	DEPLOYMENT_STRATEGY_CANARY     DeploymentStrategy = "CANARY"
	DEPLOYMENT_STRATEGY_RECREATE   DeploymentStrategy = "RECREATE"
)

const (
	PIPELINE_STAGE_TYPE_PRE_CI                       PipelineStageType                   = "PRE_CI"
	PIPELINE_STAGE_TYPE_POST_CI                      PipelineStageType                   = "POST_CI"
	PIPELINE_STAGE_TYPE_PRE_CD                       PipelineStageType                   = "PRE_CD"
	PIPELINE_STAGE_TYPE_POST_CD                      PipelineStageType                   = "POST_CD"
	PIPELINE_STEP_TYPE_INLINE                        PipelineStepType                    = "INLINE"
	PIPELINE_STEP_TYPE_REF_PLUGIN                    PipelineStepType                    = "REF_PLUGIN"
	PIPELINE_STAGE_STEP_VARIABLE_TYPE_INPUT          PipelineStageStepVariableType       = "INPUT"
	PIPELINE_STAGE_STEP_VARIABLE_TYPE_OUTPUT         PipelineStageStepVariableType       = "OUTPUT"
	PIPELINE_STAGE_STEP_VARIABLE_VALUE_TYPE_NEW      PipelineStageStepVariableValueType  = "NEW"
	PIPELINE_STAGE_STEP_VARIABLE_VALUE_TYPE_PREVIOUS PipelineStageStepVariableValueType  = "FROM_PREVIOUS_STEP"
	PIPELINE_STAGE_STEP_VARIABLE_VALUE_TYPE_GLOBAL   PipelineStageStepVariableValueType  = "GLOBAL"
	PIPELINE_STAGE_STEP_CONDITION_TYPE_SKIP          PipelineStageStepConditionType      = "SKIP"
	PIPELINE_STAGE_STEP_CONDITION_TYPE_TRIGGER       PipelineStageStepConditionType      = "TRIGGER"
	PIPELINE_STAGE_STEP_CONDITION_TYPE_SUCCESS       PipelineStageStepConditionType      = "PASS"
	PIPELINE_STAGE_STEP_CONDITION_TYPE_FAIL          PipelineStageStepConditionType      = "FAIL"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_STRING  PipelineStageStepVariableFormatType = "STRING"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_NUMBER  PipelineStageStepVariableFormatType = "NUMBER"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_BOOL    PipelineStageStepVariableFormatType = "BOOL"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_DATE    PipelineStageStepVariableFormatType = "DATE"
	SCRIPT_TYPE_CONTAINER_IMAGE                      ScriptType                          = "CONTAINER_IMAGE"
	IMAGE_PULL_TYPE_CONTAINER_REGISTRY               ScriptImagePullSecretType           = "CONTAINER_REGISTRY"
)

const (
	SEQUENTIAL CDPipelineAddType = "SEQUENTIAL"
	PARALLEL   CDPipelineAddType = "PARALLEL"
)

const (
	CD_CREATE CdPatchAction = iota
	CD_DELETE               //delete this pipeline
	CD_UPDATE
)

type CdManifestKind string

const (
	CD_PIPELINE_KIND CdManifestKind = "CD"
)
