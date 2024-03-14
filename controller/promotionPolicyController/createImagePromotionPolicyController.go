package promotionPolicyController

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/promotionPolicyClient"
)

func DownloadPolicyConfigController(payload *ArtifactPromotionPolicy.PayloadPolicyForCreate) error {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated", err)
		return err
	}
	return promotionPolicyClient.CreateImagePromotionPolicy(payload)

}
