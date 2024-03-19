package models

type CDClonePayload struct {
	SpecVersion     string      `yaml:"specVersion"` //defaulted to V1
	EnvironmentId   int         `yaml:"environmentId" validate:"required"`
	SourceEnvId     int         `yaml:"sourceEnvironmentId" validate:"required"`
	RunForALlApps   bool        `yaml:"runForALlApps"`
	ProjectIds      []int       `yaml:"projectIds"`
	Overrides       []Overrides `yaml:"overrides" json:"overrides" validate:"dive"`
	CommonOverrides Overrides   `yaml:"commonOverrides"`
	IncludesAppName string      `yaml:"includesApp"`
	ExcludesAppName string      `yaml:"excludesApp"`
}
