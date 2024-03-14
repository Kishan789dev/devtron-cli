package ArtifactPromotionPolicy

type ApplyPolicyManifest struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind" json:"kind"`
	Spec       Spec   `yaml:"spec" json:"spec"`
}

type Spec struct {
	Payload PayloadApplyPolicy `yaml:"payload" json:"payload"`
}

type PayloadApplyPolicy struct {
	ApplicationEnvironments []PolicyVariables      `yaml:"applicationEnvironments" json:"applicationEnvironments"`
	ApplyToPolicyName       string                 `yaml:"applyToPolicyName" json:"applyToPolicyName"`
	AppEnvPolicyListFilter  AppEnvPolicyListFilter `yaml:"appEnvPolicyListFilter" json:"appEnvPolicyListFilter"`
}

type PolicyVariables struct {
	AppName    string `yaml:"appName" json:"appName"`
	EnvName    string `yaml:"envName" json:"envName"`
	PolicyName string `yaml:"policyName" json:"policyName"`
}

type AppEnvPolicyListFilter struct {
	AppNames    []string `yaml:"appNames" json:"appNames"`
	EnvNames    []string `yaml:"envNames" json:"envNames"`
	PolicyNames []string `yaml:"policyNames" json:"policyNames"`
	SortBy      string   `json:"sortBy"`
	SortOrder   string   `json:"sortOrder"`
	Offset      int      `json:"offset"`
	Size        int      `json:"size"`
}
