package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

func HandleImagePromotionPolicy() {
	policyManifest, err := getDetailsForFlagInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	validate := validator.New()
	err = validate.Struct(policyManifest)
	if err != nil {
		fmt.Print("Invalid configuration", err)
		return
	}

	err = promotionPolicyController.DownloadPolicyConfigController(policyManifest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Policy Created Successfully")

	HandlerApplyPolicyHelper("applyPath")

}

func getDetailsForFlagInput() (*ArtifactPromotionPolicy.PayloadPolicyForCreate, error) {

	name := viper.GetString("name")
	description := viper.GetString("descriptionPolicy")
	pass := viper.GetString("passConditionPolicy")
	fail := viper.GetString("failConditionPolicy")
	approveCnt := viper.GetInt("approverCountPolicy")
	allowImageBuilderFromApprove := viper.GetBool("allowImageBuilderFromApprovePolicy")
	allowRequesterFromApprove := viper.GetBool("allowRequestFomApprovePolicy")
	allowApprovedDeployment := viper.GetBool("allowApproverFromDeployPolicy")
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
	if pass == "" && fail == "" {
		err := fmt.Errorf("missing condition criteria ")
		return nil, err
	}

	approvalMetadata := ArtifactPromotionPolicy.ApprovalMetadata{
		ApproverCount:                approveCnt,
		AllowImageBuilderFromApprove: allowImageBuilderFromApprove,
		AllowRequesterFromApprove:    allowRequesterFromApprove,
		AllowApproverFromDeploy:      allowApprovedDeployment,
	}

	payload := &ArtifactPromotionPolicy.PayloadPolicyForCreate{
		Name:             name,
		Description:      description,
		Conditions:       cond,
		ApprovalMetadata: approvalMetadata,
	}

	return payload, nil
}
