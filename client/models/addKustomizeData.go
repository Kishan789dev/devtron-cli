package models

type KustomizeDataPayload struct {
	AppId    int    `json:"appId"`
	EnvId    int    `json:"envId"`
	FilePath string `json:"filePath"`
}
