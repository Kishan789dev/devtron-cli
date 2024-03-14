package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
)

func CreateImagePromotionPolicy(request *ArtifactPromotionPolicy.PayloadPolicyForCreate) error {
	response := models.Response[string]{}
	err := client.CallPostApi(POLICY, request, &response)
	return err

}
