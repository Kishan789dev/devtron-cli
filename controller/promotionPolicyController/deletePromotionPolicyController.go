package promotionPolicyController

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/promotionPolicyClient"
)

func DeletePromotionPolicyController(name string) error {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return nil
	}

	if !isUserAuthenticated {
		fmt.Println("User is not authenticated")
		return nil
	}

	return promotionPolicyClient.DeletePromotionPolicyClient(name)

}
