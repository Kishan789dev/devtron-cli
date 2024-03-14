package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
)

func UpdatePolicyClient(request *ArtifactPromotionPolicy.PayloadPolicyForCreate, PolicyName string) error {
	response := models.Response[string]{}
	err := client.CallPutApi(POLICY+"/"+PolicyName, request, &response)
	return err
}
