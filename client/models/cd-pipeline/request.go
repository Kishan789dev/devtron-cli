package cd_pipeline

type CDPatchRequest struct {
	Pipeline         *CDPipelineConfigObject `json:"pipeline,omitempty"`
	AppId            int                     `json:"appId,omitempty"`
	Action           CdPatchAction           `json:"action,omitempty"`
	UserId           int32                   `json:"-"`
	ForceDelete      bool                    `json:"-"`
	NonCascadeDelete bool                    `json:"-"`
}
