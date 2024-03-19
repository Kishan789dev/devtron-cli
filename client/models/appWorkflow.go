package models

type AppWorkflow struct {
	Name      string     `json:"appName"`
	AppId     int        `json:"appId"`
	Workflows []Workflow `json:"workflows"`
}

type Workflow struct {
	AppId int    `json:"appId"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Tree  []Tree `json:"tree"`
}

type Tree struct {
	Id                         int    `json:"id,omitempty"`
	AppWorkflowId              int    `json:"appWorkflowId"`
	Type                       string `json:"type"`
	ComponentId                int    `json:"componentId"`
	ParentId                   int    `json:"parentId"`
	ParentType                 string `json:"parentType"`
	DeploymentAppDeleteRequest bool   `json:"deploymentAppDeleteRequest"`
}
