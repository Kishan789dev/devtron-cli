package createPolicy

type Conditions struct {
	ConditionType int    `json:"condtype"`
	Expression    string `json:"expression"`
	ErrorMsg      string `json:"errmsg"`
}
type ApprovalMetadata struct {
	ApproverCount                int  `json:"approvercount"`
	AllowImageBuilderFromApprove bool `json:"allowimagebuilderfromapprove"`
	AllowRequesterFromApprove    bool `json:"allowrequesterfromapprove"`
	AllowApproverFromDeploy      bool `json:"allowapproverfromdeploy"`
}

type Payload struct {
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Conditions       []Conditions     `json:"Condition"`
	ApprovalMetadata ApprovalMetadata `json:"approvalmetadata"`
}
