package ArtifactPromotionPolicy

type Conditions struct {
	ConditionType int    `json:"conditionType"`
	Expression    string `json:"expression"`
	ErrorMsg      string `json:"errorMsg"`
}
type ApprovalMetadata struct {
	ApproverCount                int  `json:"approverCount"`
	AllowImageBuilderFromApprove bool `json:"allowImageBuilderFromApprove"`
	AllowRequesterFromApprove    bool `json:"allowRequesterFromApprove"`
	AllowApproverFromDeploy      bool `json:"allowApproverFromDeploy"`
}

type PayloadPolicyForCreate struct {
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Conditions       []Conditions     `json:"Conditions"`
	ApprovalMetadata ApprovalMetadata `json:"approvalMetadata"`
}
