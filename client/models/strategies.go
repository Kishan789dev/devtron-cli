package models

type PipelineStrategy struct {
	Strategy []Strategy `json:"pipelineStrategy"`
}

type Strategy struct {
	DeploymentTemplate string `json:"deploymentTemplate"`
	Default            bool   `json:"default"`
	Config             any    `json:"config"`
}
