package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
)

func GetArtifactPromotionPolicyDetails(name string) (ArtifactPromotionPolicy.PayloadPolicyForCreate, error) {
	response := models.Response[ArtifactPromotionPolicy.PayloadPolicyForCreate]{}

	err := client.CallGetApi(POLICY+"/"+name, make(map[string]string), &response)
	return response.Result, err
}
