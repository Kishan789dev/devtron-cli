package promotionPolicyController

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/promotionPolicyClient"
)

func GetAppAndEnvListController(Payload *ArtifactPromotionPolicy.AppEnvPolicyListFilter) (ArtifactPromotionPolicy.PolicyConfig, error) {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return ArtifactPromotionPolicy.PolicyConfig{}, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return ArtifactPromotionPolicy.PolicyConfig{}, err
	}
	return promotionPolicyClient.GetAppAndEnvListing(Payload)

}
