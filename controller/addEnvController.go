package controller

import (
	"fmt"

	"github.com/devtron-labs/devtron-cli/devtctl/client"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models"
)

func AddEnv(envPayload models.EnvConfig) models.EnvConfig {
	//auth check
	isUserAuthenticated, err := client.IsUserAuthenticated()
	if err != nil {
		fmt.Printf("Auth check failed with reason: %s\n", err)
		return models.EnvConfig{}
	}

	if isUserAuthenticated != true {
		fmt.Println("user is not authenticated")
		return models.EnvConfig{}
	}
	var FailedEnvConfigs []models.EnvPayload
	if len(envPayload.EnvPayload) == 0 {
		fmt.Println("No envs found in the payload.")
		return models.EnvConfig{
			EnvPayload: FailedEnvConfigs,
		}
	}
	for _, envConfig := range envPayload.EnvPayload {
		envConfig.Active = true
		env_id, err := processEnv(envConfig)
		if err != nil {
			fmt.Printf(" unable to add environment - `%s` with namespace - `%s` in cluster-id - `%d` :  \n Reason :%s\n", envConfig.EnvironmentName, envConfig.Namespace, envConfig.ClusterId, err)
			FailedEnvConfigs = append(FailedEnvConfigs, envConfig)
		} else {
			fmt.Printf("Environment `%s` Added successfully! with id -- %d\n", envConfig.EnvironmentName, env_id)
		}
	}

	return models.EnvConfig{
		EnvPayload: FailedEnvConfigs,
	}
}
func processEnv(envConfig models.EnvPayload) (int, error) {
	id, err := client.AddEnvApiCall(envConfig)
	if err != nil {
		return 0, fmt.Errorf("error to env  `%s`  in cluster  error -- %s\n", envConfig.EnvironmentName, err)
	}
	return id, nil
}
