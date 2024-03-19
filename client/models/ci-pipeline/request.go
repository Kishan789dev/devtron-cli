package ci_pipeline

type CiPatchRequest struct {
	CiPipeline    CiPipeline  `json:"ciPipeline"`
	AppId         int         `json:"appId,omitempty"`
	Action        PatchAction `json:"action"`
	AppWorkflowId int         `json:"appWorkflowId"`
}

type PatchAction int

const (
	CREATE          PatchAction = iota
	UPDATE_PIPELINE             //update pipeline
)
