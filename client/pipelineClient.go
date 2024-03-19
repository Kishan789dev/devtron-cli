package client

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func AddKustomizeDataInZip(appId int, envId int, zipPath string) error {
	path := fmt.Sprintf(ADD_KUSTOMIZE_DATA, appId, envId)
	response := models.Response[string]{}
	err := CallPostApiWithFiles(path, map[string]string{"file": zipPath}, map[string]string{}, &response)
	if err != nil {
		return err
	}
	fmt.Println("Success")
	return nil
}
