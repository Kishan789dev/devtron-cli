package client

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/createPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/getListOfPolicies"
)

func GetListOfPolicies(params *getListOfPolicies.PoliciesList) ([]createPolicy.Payload, error) {
	//var response []createPolicy.Payload
	response := models.Response[[]createPolicy.Payload]{}

	query := map[string]string{"search": params.Search, "sortby": params.SortBy, "sortorder": params.SortOrder}
	err := CallGetApi(POLICY, query, &response)
	fmt.Println(err)
	fmt.Println(response)
	return response.Result, err

}
