package promotionPolicyClient

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func DeletePromotionPolicyClient(name string) error {
	response := models.Response[string]{}

	err := client.CallDeleteApi(POLICY+"/"+name, make(map[string]string), &response)
	return err

}
