package azure

import (
	"github.com/openshift/installer/pkg/terraform"
	"github.com/openshift/installer/pkg/terraform/stages"
)

const azure string = "azure"
const azurestack string = "azurestack"

// PlatformStages are the stages to run to provision the infrastructure in Azure.
var PlatformStages = []terraform.Stage{
	stages.NewStage(azure, "vnet"),
	stages.NewStage(azure, "bootstrap", stages.WithNormalDestroy()),
	stages.NewStage(azure, "cluster"),
}

// StackPlatformStages are the stages to run to provision the infrastructure in Azure Stack.
var StackPlatformStages = []terraform.Stage{
	stages.NewStage(azurestack, "vnet"),
	stages.NewStage(azurestack, "bootstrap", stages.WithNormalDestroy()),
	stages.NewStage(azurestack, "cluster"),
}
