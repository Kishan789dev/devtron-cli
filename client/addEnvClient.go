package client

import (
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func AddEnvApiCall(request models.EnvPayload) (int, error) {
	response := models.Response[models.EnvPayload]{}
	err := CallPostApi(ADD_ENVIRONMENT, request, &response)
	return response.Result.Id, err
}
