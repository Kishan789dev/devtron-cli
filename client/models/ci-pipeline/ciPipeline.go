package ci_pipeline

type CPipelineManifest struct {
	ApiVersion string   `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string   `yaml:"kind" json:"kind" validate:"required"`
	Metadata   Metadata `yaml:"metadata" json:"metadata" validate:"required" validate:"dive"`
	Spec       Spec     `yaml:"spec" json:"spec" validate:"required" validate:"dive"`
}
type Metadata struct {
	Type ManifestType `yaml:"type" json:"type" validate:"required"`
}

type Spec struct {
	Payload []Payload `yaml:"payload" json:"payload" validate:"required" validate:"dive"`
}

type Payload struct {
	Criteria Criteria `yaml:"criteria" json:"criteria" validate:"dive"`
}
type Criteria struct {
	PipelineIds []int `yaml:"pipelineIds" json:"pipelineIds"`
	AppIds      []int `yaml:"appIds" json:"appIds"`
	//IncludesPipelineNames []string `yaml:"includesPipelineNames" json:"includesPipelineNames"`
	//ExcludesPipelineNames []string `yaml:"excludesPipelineNames" json:"excludesPipelineNames"`
	IncludesAppNames  []string `yaml:"includesAppNames" json:"includesAppNames"`
	ExcludesAppNames  []string `yaml:"excludesAppNames" json:"excludesAppNames"`
	EnvironmentNames  []string `yaml:"environmentNames" json:"environmentNames"`
	ProjectNames      []string `yaml:"projectNames" json:"projectNames"`
	AppendPreCiSteps  bool     `yaml:"appendPreCiSteps" json:"appendPreCiSteps"`
	AppendPostCiSteps bool     `yaml:"appendPostCiSteps" json:"appendPostCiSteps"`
}
