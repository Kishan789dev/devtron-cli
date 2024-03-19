package models

type EnvConfig struct {
	EnvPayload []EnvPayload `yaml:"envPayload"`
}

type EnvPayload struct {
	Id              int    `yaml:"id" json:"id"`
	ClusterId       int    `yaml:"cluster_id" json:"cluster_id"`
	EnvironmentName string `yaml:"environment_name"  json:"environment_name"`
	Namespace       string `yaml:"namespace" json:"namespace"`
	Active          bool   `yaml:"active" json:"active"`
	IsProduction    bool   `yaml:"is_production" json:"default"`
	Description     string `yaml:"description" json:"description"`
}
