package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
)

func GetListOfPolicies(params *ArtifactPromotionPolicy.PoliciesList) ([]ArtifactPromotionPolicy.PayloadPolicyForCreate, error) {
	response := models.Response[[]ArtifactPromotionPolicy.PayloadPolicyForCreate]{}

	query := map[string]string{"search": params.Search, "sortBy": params.SortBy, "sortOrder": params.SortOrder}
	err := client.CallGetApi(POLICY, query, &response)
	return response.Result, err

}
