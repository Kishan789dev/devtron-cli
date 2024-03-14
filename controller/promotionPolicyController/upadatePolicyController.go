package promotionPolicyController

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/promotionPolicyClient"
)

func UpdatePolicyController(payload *ArtifactPromotionPolicy.PayloadPolicyForCreate, PolicyName string) error {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return err
	}
	return promotionPolicyClient.UpdatePolicyClient(payload, PolicyName)

}
