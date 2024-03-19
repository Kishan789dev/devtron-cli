package handler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/createPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller"

	//"github.com/devtron-labs/devtron-cli/devtctl/controller"

	"github.com/spf13/viper"
)

func HandleImagePromotionPolicy() {
	policymanifest, err := getDetailsForFlagInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = controller.DownloadPolicyConfigController(policymanifest)
	if err != nil {
		fmt.Println(err)
	}
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//if response_status == 200 {
	//	fmt.Println("policy created successfully")
	//	return
	//}
	//
	//if response_status == 400 {
	//	fmt.Println("this response is for any un acceptable request payload or query params")
	//	return
	//}
	//
	//if response_status == 403 {
	//	fmt.Println("gets this response if user is not a super admin")
	//	return
	//}

}
func getDetailsForFlagInput() (*createPolicy.Payload, error) {

	name := viper.GetString("name")
	description := viper.GetString("description")
	pass := viper.GetString("pass_condition")
	fail := viper.GetString("fail_condition")
	approverCnt := viper.GetInt("approver_count")
	allowImageBuilderFromAppr := viper.GetBool("allow_image_builder_from_approve")
	allowRequesterFromApprove := viper.GetBool("allow_request_fom_approve")
	allowApproverFromDep := viper.GetBool("allow_approver_from_deploy")

	//cond1 := &createPolicy.Conditions{}
	var cond []createPolicy.Conditions

	//is the name flag arbitrary ??

	if pass != "" {
		cond1 := createPolicy.Conditions{ConditionType: 1,
			Expression: pass,
			ErrorMsg:   "nil",
		}
		cond = append(cond, cond1)

	}

	if fail != "" {
		cond1 := createPolicy.Conditions{ConditionType: 0,
			Expression: fail,
			ErrorMsg:   "nil",
		}
		cond = append(cond, cond1)

	}
	if pass == "" && fail == "" {
		err := fmt.Errorf("missing condition criteria ")
		return nil, err
	}

	approvalmetadata := createPolicy.ApprovalMetadata{
		ApproverCount:                approverCnt,
		AllowImageBuilderFromApprove: allowImageBuilderFromAppr,
		AllowRequesterFromApprove:    allowRequesterFromApprove,
		AllowApproverFromDeploy:      allowApproverFromDep,
	}

	payload := &createPolicy.Payload{
		Name:             name,
		Description:      description,
		Conditions:       cond,
		ApprovalMetadata: approvalmetadata,
	}
	return payload, nil
}
