package applyPolicy

type ApplyPolicyManifest struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" json:"kind" validate:"required"`
	Spec       Spec   `yaml:"spec" json:"spec" validate:"required" validate:"dive"`
}

type Spec struct {
	Payload []Payload `yaml:"payload" json:"payload" validate:"required" validate:"dive"`
}
type Payload struct {
	ApplicationEnvironments []PolicyVariables `json:"applicationEnvironments"`

	ApplyToPolicyName      string                 `json:"applyToPolicyName"`
	AppEnvPolicyListFilter AppEnvPolicyListFilter `json:"appEnvPolicyListFilter"`
}

type PolicyVariables struct {
	AppName    string `json:"appName"`
	EnvName    string `json:"envName"`
	PolicyName string `json:"policyName"`
}
type AppEnvPolicyListFilter struct {
	AppNames    []string `json:"appNames"`
	EnvNames    []string `json:"envNames"`
	PolicyNames []string `json:"policyNames"`
	//SortBy      string   `json:"sortBy"`
	//SortOrder   string   `json:"sortOrder"`
	//Offset      int      `json:"offset"`
	//Size        int      `json:"size"`
}
