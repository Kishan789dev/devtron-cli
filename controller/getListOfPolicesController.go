package controller

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/createPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/getListOfPolicies"
)

func GetPoliciesList(paramsdetails *getListOfPolicies.PoliciesList) ([]createPolicy.Payload, error) {

	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		err := fmt.Errorf("Auth check failed with reason:")
		return []createPolicy.Payload{}, err
	}

	if isUserAuthenticated != true {
		fmt.Println("User is not authenticated")
		return []createPolicy.Payload{}, err
	}
	return client.GetListOfPolicies(paramsdetails)

}
