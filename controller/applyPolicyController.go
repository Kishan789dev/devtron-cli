package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/applyPolicy"
)

func ApplyPolicyController(ciPayload applyPolicy.ApplyPolicyManifest) (applyPolicy.ApplyPolicyManifest, error) {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return applyPolicy.ApplyPolicyManifest{}, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return applyPolicy.ApplyPolicyManifest{}, err
	}
	//
	//applyPolicyManifest, err := getPrePostConfigManifest(ciPayload)
	//if err != nil {
	//	return applyPolicy.ApplyPolicyManifest{}, err
	//}
	//
	return applyPolicy.ApplyPolicyManifest{}, nil

}
