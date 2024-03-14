package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
	"github.com/spf13/viper"
)

func HandleGetArtifactPromotionPolicy() {
	name := viper.GetString("policyName")
	response, err := promotionPolicyController.GetArtifactPromotionPolicyController(name)
	if err != nil {

		println(err)
		return
	}
	err = utils.WriteOutputToFileInJson(response)
	if err != nil {
		fmt.Println("Error occurred during writing to json")
		return
	}
}
