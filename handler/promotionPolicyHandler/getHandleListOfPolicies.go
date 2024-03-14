package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

func HandleListOfPolicies() {
	payloadForListOfPolices := flagInputForListOfPolicies()

	validate := validator.New()
	err := validate.Struct(payloadForListOfPolices)
	if err != nil {
		fmt.Print("Invalid configuration", err)
		return
	}

	response, err := promotionPolicyController.GetPoliciesList(payloadForListOfPolices)
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
func flagInputForListOfPolicies() *ArtifactPromotionPolicy.PoliciesList {
	search := viper.GetString("searchPolicyList")
	sortBy := viper.GetString("sortByPolicyList")
	sortOrder := viper.GetString("sortOrderPolicyList")

	ParamsManifest := &ArtifactPromotionPolicy.PoliciesList{
		Search:    search,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
	return ParamsManifest

}
