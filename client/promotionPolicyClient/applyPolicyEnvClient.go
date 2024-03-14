package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	//cipipeline "github.com/devtron-labs/devtron-cli/devtctl/client/models/ci-pipeline"
)

func ApplyPolicyEnv(request ArtifactPromotionPolicy.ApplyPolicyManifest) error {

	response := models.Response[string]{}
	err := client.CallPostApi(APPLY_POLICY, request.Spec.Payload, &response)
	return err

}
