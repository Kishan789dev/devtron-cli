package ci_pipeline

type PluginType string
type ScriptType string
type ScriptImagePullSecretType string
type ScriptMappingType string
type PluginStepType string
type PluginStepVariableType string
type PluginStepVariableValueType string
type PluginStepConditionType string
type PluginStepVariableFormatType string

const (
	PLUGIN_TYPE_SHARED                  PluginType                   = "SHARED"
	PLUGIN_TYPE_PRESET                  PluginType                   = "PRESET"
	SCRIPT_TYPE_SHELL                   ScriptType                   = "SHELL"
	SCRIPT_TYPE_DOCKERFILE              ScriptType                   = "DOCKERFILE"
	SCRIPT_TYPE_CONTAINER_IMAGE         ScriptType                   = "CONTAINER_IMAGE"
	IMAGE_PULL_TYPE_CONTAINER_REGISTRY  ScriptImagePullSecretType    = "CONTAINER_REGISTRY"
	IMAGE_PULL_TYPE_SECRET_PATH         ScriptImagePullSecretType    = "SECRET_PATH"
	SCRIPT_MAPPING_TYPE_FILE_PATH       ScriptMappingType            = "FILE_PATH"
	SCRIPT_MAPPING_TYPE_DOCKER_ARG      ScriptMappingType            = "DOCKER_ARG"
	SCRIPT_MAPPING_TYPE_PORT            ScriptMappingType            = "PORT"
	PLUGIN_STEP_TYPE_INLINE             PluginStepType               = "INLINE"
	PLUGIN_STEP_TYPE_REF_PLUGIN         PluginStepType               = "REF_PLUGIN"
	PLUGIN_VARIABLE_TYPE_INPUT          PluginStepVariableType       = "INPUT"
	PLUGIN_VARIABLE_TYPE_OUTPUT         PluginStepVariableType       = "OUTPUT"
	PLUGIN_VARIABLE_VALUE_TYPE_NEW      PluginStepVariableValueType  = "NEW"
	PLUGIN_VARIABLE_VALUE_TYPE_PREVIOUS PluginStepVariableValueType  = "FROM_PREVIOUS_STEP"
	PLUGIN_VARIABLE_VALUE_TYPE_GLOBAL   PluginStepVariableValueType  = "GLOBAL"
	PLUGIN_CONDITION_TYPE_SKIP          PluginStepConditionType      = "SKIP"
	PLUGIN_CONDITION_TYPE_TRIGGER       PluginStepConditionType      = "TRIGGER"
	PLUGIN_CONDITION_TYPE_SUCCESS       PluginStepConditionType      = "SUCCESS"
	PLUGIN_CONDITION_TYPE_FAIL          PluginStepConditionType      = "FAIL"
	PLUGIN_VARIABLE_FORMAT_TYPE_STRING  PluginStepVariableFormatType = "STRING"
	PLUGIN_VARIABLE_FORMAT_TYPE_NUMBER  PluginStepVariableFormatType = "NUMBER"
	PLUGIN_VARIABLE_FORMAT_TYPE_BOOL    PluginStepVariableFormatType = "BOOL"
	PLUGIN_VARIABLE_FORMAT_TYPE_DATE    PluginStepVariableFormatType = "DATE"
)

type PipelineStageType string
type PipelineStepType string
type PipelineStageStepVariableType string
type PipelineStageStepVariableValueType string
type PipelineStageStepConditionType string
type PipelineStageStepVariableFormatType string

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
	PIPELINE_STAGE_STEP_CONDITION_TYPE_SUCCESS       PipelineStageStepConditionType      = "SUCCESS"
	PIPELINE_STAGE_STEP_CONDITION_TYPE_FAIL          PipelineStageStepConditionType      = "FAIL"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_STRING  PipelineStageStepVariableFormatType = "STRING"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_NUMBER  PipelineStageStepVariableFormatType = "NUMBER"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_BOOL    PipelineStageStepVariableFormatType = "BOOL"
	PIPELINE_STAGE_STEP_VARIABLE_FORMAT_TYPE_DATE    PipelineStageStepVariableFormatType = "DATE"
)

type TriggerType string //HOW pipeline should be triggered

const TRIGGER_TYPE_AUTOMATIC TriggerType = "AUTOMATIC"
const TRIGGER_TYPE_MANUAL TriggerType = "MANUAL"

type SourceType string

const (
	SOURCE_TYPE_BRANCH_FIXED SourceType = "SOURCE_TYPE_BRANCH_FIXED"
	SOURCE_TYPE_BRANCH_REGEX SourceType = "SOURCE_TYPE_BRANCH_REGEX"
	SOURCE_TYPE_TAG_ANY      SourceType = "SOURCE_TYPE_TAG_ANY"
	SOURCE_TYPE_WEBHOOK      SourceType = "WEBHOOK"
)

type ManifestType string

const (
	JSON_DOWNLOAD ManifestType = "JSON_DOWNLOAD"
	YAML_DOWNLOAD ManifestType = "YAML_DOWNLOAD"
	PATCH         ManifestType = "PATCH"
)

const VERSION_V1 = "v1"

type CiManifestKind string

const (
	CI_PIPELINE_KIND CiManifestKind = "CI"
)

type ArtifactPromotionPolicy string

const (
	ARTIFACT_PROMOTION_POLICY ArtifactPromotionPolicy = "artifactPromotionPolicy"
)
