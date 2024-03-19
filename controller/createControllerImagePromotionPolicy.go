package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/createPolicy"
)

func DownloadPolicyConfigController(payload *createPolicy.Payload) error {
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		err := fmt.Errorf("Auth check failed with reason:")
		return err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return err
	}
	return client.CreateImagePromotionPolicy(payload)

}
