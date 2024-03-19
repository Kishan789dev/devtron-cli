package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/getListOfPolicies"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"
	"github.com/spf13/viper"
)

func HandleListOfPolicies() {
	search := viper.GetString("search")
	sortby := viper.GetString("sortby")
	sortorder := viper.GetString("sortorder")

	Paramsmanifest := &getListOfPolicies.PoliciesList{
		Search:    search,
		SortBy:    sortby,
		SortOrder: sortorder,
	}
	Response, err := controller.GetPoliciesList(Paramsmanifest)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, Policy := range Response {
		fmt.Printf("name :%s\n", Policy.Name)
		fmt.Printf("description :%s\n", Policy.Description)
		//conditions
		for _, codition := range Policy.Conditions {
			fmt.Printf("conditionType :%d\n", codition.ConditionType)
			fmt.Printf("expression :%s\n", codition.Expression)
			fmt.Printf("errorMsg :%s\n", codition.ErrorMsg)
		}
		//ApprovalMetadata
		//for _, Approve := range Policy {
		fmt.Printf("approverCount:%d\n", Policy.ApprovalMetadata.ApproverCount)
		fmt.Printf("allowImageBuilderFromApprove:%t\n", Policy.ApprovalMetadata.AllowImageBuilderFromApprove)
		fmt.Printf("allowRequesterFromApprove:%t\n", Policy.ApprovalMetadata.AllowRequesterFromApprove)
		fmt.Printf("allowApproverFromDeploy:%t\n", Policy.ApprovalMetadata.AllowApproverFromDeploy)

	}

	//policymanifest, err := gettingDetailsForFlagInput()
}

//
//func gettingDetailsForFlagInput() {
//
//}
