package client

import "github.com/devtron-labs/devtron-cli/devtctl/client/models"

func IsUserAuthenticated() (bool, error) {
	response := models.Response[bool]{}
	err := CallGetApi(USER_AUTH, make(map[string]string), &response)
	return response.Result, err
}
