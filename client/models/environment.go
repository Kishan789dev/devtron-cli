package models

type Environment struct {
	Name                  string     `json:"environment_name"`
	Id                    int        `json:"id"`
	Workflows             []Workflow `json:"workflows"`
	ClusterId             int        `json:"cluster_id"`
	Active                bool       `json:"active"`
	Default               bool       `json:"default"`
	PrometheusEndpoint    string     `json:"prometheusEndpoint"`
	Namespace             string     `json:"namespace"`
	IsClusterCdActive     bool       `json:"isClusterCdActive"`
	EnvironmentIdentifier string     `json:"environmentIdentifier"`
	AppCount              int        `json:"appCount"`
}
