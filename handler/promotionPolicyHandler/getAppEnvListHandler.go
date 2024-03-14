package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

func HandleGetAppAndEnvList() {
	payloadForAppEnvList, err := flapInputForAppEnvList()

	if err != nil {
		fmt.Println(err)
		return
	}

	validate := validator.New()
	err = validate.Struct(payloadForAppEnvList)
	if err != nil {
		fmt.Print("Invalid configuration", err)
		return
	}

	response, err := promotionPolicyController.GetAppAndEnvListController(payloadForAppEnvList)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = utils.WriteOutputToFileInJson(response)
	if err != nil {
		fmt.Println("Error occurred during writing to json")
		return
	}

}
func flapInputForAppEnvList() (*ArtifactPromotionPolicy.AppEnvPolicyListFilter, error) {
	appNames := utils.SplitAndTrim(viper.GetString("appNamesList"))
	envNames := utils.SplitAndTrim(viper.GetString("envNamesList"))
	policyNames := utils.SplitAndTrim(viper.GetString("policyNamesList"))
	sortBy := viper.GetString("sortBy")
	sortOrder := viper.GetString("sortOrder")
	offsetStr := viper.GetString("offset")
	sizeStr := viper.GetString("size")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		fmt.Println("Error converting offsetStr to offset:", err)
		return nil, err
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		fmt.Println("Error converting size to integer:", err)
		return nil, err
	}

	GetAppEnvListPayload := &ArtifactPromotionPolicy.AppEnvPolicyListFilter{
		AppNames:    appNames,
		EnvNames:    envNames,
		PolicyNames: policyNames,
		SortBy:      sortBy,
		SortOrder:   sortOrder,
		Offset:      offset,
		Size:        size,
	}
	return GetAppEnvListPayload, nil
}
