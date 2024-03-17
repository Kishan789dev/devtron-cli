package ArtifactPromotionPolicy

type ApplyPolicyManifest struct {
	ApiVersion string `yaml:"apiVersion" json:"apiVersion" validate:"required"`
	Kind       string `yaml:"kind" json:"kind" validate:"required"`
	Spec       Spec   `yaml:"spec" json:"spec" validate:"dive"`
}

type Spec struct {
	Payload ApplyPolicy `yaml:"payload" json:"payload" validate:"dive"`
}

type ApplyPolicy struct {
	ApplicationEnvironments []PolicyVariables      `yaml:"applicationEnvironments" json:"applicationEnvironments" validate:"dive"`
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
