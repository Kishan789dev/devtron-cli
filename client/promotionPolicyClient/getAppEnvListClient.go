package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
)

func GetAppAndEnvListing(request *ArtifactPromotionPolicy.AppEnvPolicyListFilter) (ArtifactPromotionPolicy.PolicyConfig, error) {

	response := models.Response[ArtifactPromotionPolicy.PolicyConfig]{}
	err := client.CallPostApi(APP_ENV_LIST, request, &response)
	return response.Result, err

}
