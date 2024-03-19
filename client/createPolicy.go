package client

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/createPolicy"
)

func CreateImagePromotionPolicy(request *createPolicy.Payload) error {
	response := models.Response[any]{}
	err := CallPostApi(POLICY, request, &response)

	return err

}
