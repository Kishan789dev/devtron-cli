package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/jedib0t/go-pretty/v6/table"
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

	t1 := table.NewWriter()
	t1.SetStyle(table.StyleBold)
	if payloadForListOfPolices.Expand == true {

		t1.AppendHeader(table.Row{"Name", "Description", "Pass Condition", "Fail_Condition", "ApproverCount", "ImageBuilderApprove", "RequesterApprove", "ApproverDeploy"})
		for _, list := range response {
			if len(list.Conditions) > 1 {
				if list.Conditions[0].ConditionType == 1 {
					t1.AppendRow(table.Row{list.Name, list.Description, list.Conditions[0].Expression, list.Conditions[1].Expression, list.ApprovalMetadata.ApproverCount, list.ApprovalMetadata.AllowImageBuilderFromApprove, list.ApprovalMetadata.AllowRequesterFromApprove, list.ApprovalMetadata.AllowApproverFromDeploy})

				} else {
					t1.AppendRow(table.Row{list.Name, list.Description, list.Conditions[1].Expression, list.Conditions[0].Expression, list.ApprovalMetadata.ApproverCount, list.ApprovalMetadata.AllowImageBuilderFromApprove, list.ApprovalMetadata.AllowRequesterFromApprove, list.ApprovalMetadata.AllowApproverFromDeploy})
				}
			} else {
				if list.Conditions[0].ConditionType == 1 {
					t1.AppendRow(table.Row{list.Name, list.Description, list.Conditions[0].Expression, "", list.ApprovalMetadata.ApproverCount, list.ApprovalMetadata.AllowImageBuilderFromApprove, list.ApprovalMetadata.AllowRequesterFromApprove, list.ApprovalMetadata.AllowApproverFromDeploy})

				} else {
					t1.AppendRow(table.Row{list.Name, list.Description, "", list.Conditions[0].Expression, list.ApprovalMetadata.ApproverCount, list.ApprovalMetadata.AllowImageBuilderFromApprove, list.ApprovalMetadata.AllowRequesterFromApprove, list.ApprovalMetadata.AllowApproverFromDeploy})
				}
			}
		}

	} else {
		t1.AppendHeader(table.Row{"Name", "Description"})

		for _, list := range response {

			t1.AppendRow(table.Row{list.Name, list.Description})

		}

	}
	fmt.Println(t1.Render())

}
func flagInputForListOfPolicies() *ArtifactPromotionPolicy.PoliciesList {
	search := viper.GetString("searchPolicyList")
	sortBy := viper.GetString("sortByPolicyList")
	sortOrder := viper.GetString("sortOrderPolicyList")
	expand := viper.GetBool("expand")

	ParamsManifest := &ArtifactPromotionPolicy.PoliciesList{
		Search:    search,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Expand:    expand,
	}
	return ParamsManifest

}
