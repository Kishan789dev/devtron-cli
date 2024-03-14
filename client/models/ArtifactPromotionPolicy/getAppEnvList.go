package ArtifactPromotionPolicy

type PolicyConfig struct {
	TotalCount                   int         `json:"totalCount"`
	AppEnvironmentPolicyMappings []PolicyMap `json:"appEnvironmentPolicyMappings"`
}
type PolicyMap struct {
	AppName    string `json:"appName"`
	EnvName    string `json:"envName"`
	PolicyName string `json:"policyName"`
}
