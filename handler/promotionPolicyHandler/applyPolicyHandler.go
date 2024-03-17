package promotionPolicyHandler

import (
	"fmt"
	"github.com/devtron-labs/devtron-cli/devtctl/client/models/ArtifactPromotionPolicy"
	"github.com/devtron-labs/devtron-cli/devtctl/controller/promotionPolicyController"
	"github.com/devtron-labs/devtron-cli/devtctl/handler/utils"
)

func HandleApplyPolicy() {
	HandlerApplyPolicyHelper("path")
}

func HandlerApplyPolicyHelper(path string) {
	var applyPolicyManifest ArtifactPromotionPolicy.ApplyPolicyManifest

	applyPolicyManifest, err := utils.ReadInputFile(applyPolicyManifest, path)
	if err != nil {
		fmt.Print("Bad input file", err)
		return
	}

	if applyPolicyManifest.ApiVersion != ArtifactPromotionPolicy.VERSION_v1 {
		fmt.Println("Invalid version provided in manifest, please use v1 ")
		return
	}

	if applyPolicyManifest.Kind != string(ArtifactPromotionPolicy.ARTIFACT_PROMOTION_POLICY) {
		fmt.Println("Invalid kind provided in manifest, did you mean CI-Pipeline ")
		return
	}

	err = promotionPolicyController.ApplyPolicyController(applyPolicyManifest)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Policy Applied Successfully")
}
