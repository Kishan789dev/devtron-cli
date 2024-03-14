package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/spf13/viper"
)

func HandledeletePromotionPolicy() {
	name := viper.GetString("policyNameDelete")
	err := promotionPolicyController.DeletePromotionPolicyController(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Deleted Successfully")
}
