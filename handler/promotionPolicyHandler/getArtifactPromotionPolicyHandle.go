package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
)

func HandleGetArtifactPromotionPolicy() {
	name := viper.GetString("policyNameOfArtifactPromotionPolicy")
	response, err := promotionPolicyController.GetArtifactPromotionPolicyController(name)
	if err != nil {
		println(err)
		return
	}

	t1 := table.NewWriter()
	t1.SetStyle(table.StyleBold)

	t1.AppendRow(table.Row{"Name", response.Name})
	t1.AppendRow(table.Row{"Description", response.Description})
	if len(response.Conditions) > 1 {
		if response.Conditions[0].ConditionType == 1 {
			t1.AppendRow(table.Row{"Pass Condition", response.Conditions[0].Expression})
			t1.AppendRow(table.Row{"Fail Condition", response.Conditions[1].Expression})
		} else {

			t1.AppendRow(table.Row{"Pass Condition", response.Conditions[1].Expression})
			t1.AppendRow(table.Row{"Fail Condition", response.Conditions[0].Expression})
		}

	} else {

		if response.Conditions[0].ConditionType == 1 {
			t1.AppendRow(table.Row{"Pass Condition", response.Conditions[0].Expression})
		} else {

			t1.AppendRow(table.Row{"Fail Condition", response.Conditions[0].Expression})
		}
	}
	t1.AppendRow(table.Row{"ApproverCount", response.ApprovalMetadata.ApproverCount})
	t1.AppendRow(table.Row{"AllowImageBuilderFromApprove", response.ApprovalMetadata.AllowImageBuilderFromApprove})
	t1.AppendRow(table.Row{"AllowRequesterFromApprove", response.ApprovalMetadata.AllowRequesterFromApprove})
	t1.AppendRow(table.Row{"AllowApproverFromDeploy", response.ApprovalMetadata.AllowApproverFromDeploy})
	fmt.Println(t1.Render())
	//
	//err = utils.WriteOutputToFileInJson(response)
	//if err != nil {
	//	fmt.Println("Error occurred during writing to json")
	//	return
	//}
}
