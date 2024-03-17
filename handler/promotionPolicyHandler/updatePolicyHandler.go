package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

func HandleUpdatePolicy() {

	payloadForUpdate, PolicyName, err := CreatingPayload()
	if err != nil {
		fmt.Println(err)
		return
	}

	validate := validator.New()
	err = validate.Struct(payloadForUpdate)
	if err != nil {
		fmt.Print("Invalid configuration", err)
		return
	}

	err = promotionPolicyController.UpdatePolicyController(payloadForUpdate, PolicyName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Updated Successfully")

}

func CreatingPayload() (*ArtifactPromotionPolicy.PayloadPolicyForCreate, string, error) {

	oldName := viper.GetString("policyName")
	newName := viper.GetString("updatedPolicyName")

	description := viper.GetString("descriptionUpdate")
	pass := viper.GetString("passConditionUpdate")
	fail := viper.GetString("failConditionUpdate")
	approveCnt := viper.GetInt("approverCountUpdate")
	allowImageBuilderFromAppr := viper.GetBool("allowImageBuilderFromApproveUpdate")
	allowRequesterFromApprove := viper.GetBool("allowRequestFromApproveUpdate")
	allowApproverFromDep := viper.GetBool("allowApproverFromDeployUpdate")

	var cond []ArtifactPromotionPolicy.Conditions

	if pass != "" {
		cond1 := ArtifactPromotionPolicy.Conditions{ConditionType: 1,
			Expression: pass,
			ErrorMsg:   "nil",
		}
		cond = append(cond, cond1)

	}

	if fail != "" {
		cond1 := ArtifactPromotionPolicy.Conditions{ConditionType: 0,
			Expression: fail,
			ErrorMsg:   "nil",
		}
		cond = append(cond, cond1)

	}
	//if pass == "" && fail == "" {
	//	err := fmt.Errorf("missing condition criteria ")
	//	return &ArtifactPromotionPolicy.PayloadPolicyForCreate{}, "", err
	//}

	approvalmetadata := ArtifactPromotionPolicy.ApprovalMetadata{
		ApproverCount:                approveCnt,
		AllowImageBuilderFromApprove: allowImageBuilderFromAppr,
		AllowRequesterFromApprove:    allowRequesterFromApprove,
		AllowApproverFromDeploy:      allowApproverFromDep,
	}

	payload := &ArtifactPromotionPolicy.PayloadPolicyForCreate{
		Name:             newName,
		Description:      description,
		Conditions:       cond,
		ApprovalMetadata: approvalmetadata,
	}
	return payload, oldName, nil

}
