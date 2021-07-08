package ibmcloud

import (
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/core"
	"github.com/IBM/networking-go-sdk/dnsrecordsv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/golang/mock/gomock"

	"github.com/openshift/installer/pkg/asset/installconfig/ibmcloud/mock"
	"github.com/openshift/installer/pkg/ipnet"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/ibmcloud"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type editFunctions []func(ic *types.InstallConfig)

var (
	validRegion                  = "us-south"
	validCIDR                    = "10.0.0.0/16"
	validCISInstanceCRN          = "crn:v1:bluemix:public:internet-svcs:global:a/valid-account-id:valid-instance-id::"
	validClusterName             = "valid-cluster-name"
	validDNSZoneID               = "valid-zone-id"
	validBaseDomain              = "valid.base.domain"
	validVPC                     = "valid-vpc"
	validPublicSubnetUSSouth1ID  = "public-subnet-us-south-1-id"
	validPublicSubnetUSSouth2ID  = "public-subnet-us-south-2-id"
	validPrivateSubnetUSSouth1ID = "private-subnet-us-south-1-id"
	validPrivateSubnetUSSouth2ID = "private-subnet-us-south-2-id"
	validSubnets                 = []string{
		validPublicSubnetUSSouth1ID,
		validPublicSubnetUSSouth2ID,
		validPrivateSubnetUSSouth1ID,
		validPrivateSubnetUSSouth2ID,
	}
	validZoneUSSouth1 = "us-south-1"

	validInstanceProfies = []vpcv1.InstanceProfile{{Name: &[]string{"type-a"}[0]}, {Name: &[]string{"type-b"}[0]}}

	validVPCConfig = func(ic *types.InstallConfig) {
		ic.IBMCloud.VPC = validVPC
		ic.IBMCloud.Subnets = validSubnets
	}
	notFoundVPC            = func(ic *types.InstallConfig) { ic.IBMCloud.VPC = "not-found" }
	internalErrorVPC       = func(ic *types.InstallConfig) { ic.IBMCloud.VPC = "internal-error-vpc" }
	subnetInvalidZone      = func(ic *types.InstallConfig) { ic.IBMCloud.Subnets = []string{"subnet-invalid-zone"} }
	machinePoolInvalidType = func(ic *types.InstallConfig) {
		ic.ControlPlane.Platform.IBMCloud = &ibmcloud.MachinePool{
			InstanceType: "invalid-type",
		}
	}

	existingDNSRecordsResponse = []dnsrecordsv1.DnsrecordDetails{
		{
			ID: core.StringPtr("valid-dns-record-1"),
		},
		{
			ID: core.StringPtr("valid-dns-record-2"),
		},
	}
	noDNSRecordsResponse = []dnsrecordsv1.DnsrecordDetails{}
)

func validInstallConfig() *types.InstallConfig {
	return &types.InstallConfig{
		ObjectMeta: metav1.ObjectMeta{
			Name: validClusterName,
		},
		BaseDomain: validBaseDomain,
		Networking: &types.Networking{
			MachineNetwork: []types.MachineNetworkEntry{
				{CIDR: *ipnet.MustParseCIDR(validCIDR)},
			},
		},
		Publish: types.ExternalPublishingStrategy,
		Platform: types.Platform{
			IBMCloud: validMinimalPlatform(),
		},
		ControlPlane: &types.MachinePool{
			Platform: types.MachinePoolPlatform{
				IBMCloud: validMachinePool(),
			},
		},
		Compute: []types.MachinePool{{
			Platform: types.MachinePoolPlatform{
				IBMCloud: validMachinePool(),
			},
		}},
	}
}

func validMinimalPlatform() *ibmcloud.Platform {
	return &ibmcloud.Platform{
		Region: validRegion,
	}
}

func validMachinePool() *ibmcloud.MachinePool {
	return &ibmcloud.MachinePool{}
}

func TestValidate(t *testing.T) {
	cases := []struct {
		name     string
		edits    editFunctions
		errorMsg string
	}{
		{
			name:     "valid install config",
			edits:    editFunctions{},
			errorMsg: "",
		},
		{
			name:     "valid vpc config",
			edits:    editFunctions{validVPCConfig},
			errorMsg: "",
		},
		{
			name:     "not found vpc",
			edits:    editFunctions{validVPCConfig, notFoundVPC},
			errorMsg: `^platform\.ibmcloud\.vpc: Not found: \"not-found\"$`,
		},
		{
			name:     "internal error vpc",
			edits:    editFunctions{validVPCConfig, internalErrorVPC},
			errorMsg: `^platform\.ibmcloud\.vpc: Internal error$`,
		},
		{
			name:     "subnet invalid zone",
			edits:    editFunctions{validVPCConfig, subnetInvalidZone},
			errorMsg: `^\Qplatform.ibmcloud.subnets[0]: Invalid value: "subnet-invalid-zone": subnet is not in expected zones: [us-south-1 us-south-2 us-south-3]\E$`,
		},
		{
			name:     "machine pool invalid type",
			edits:    editFunctions{validVPCConfig, machinePoolInvalidType},
			errorMsg: `^\QcontrolPlane.platform.ibmcloud.type: Not found: "invalid-type"\E$`,
		},
		{
			name:     "machine pool invalid type",
			edits:    editFunctions{validVPCConfig, machinePoolInvalidType},
			errorMsg: `^\QcontrolPlane.platform.ibmcloud.type: Not found: "invalid-type"\E$`,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ibmcloudClient := mock.NewMockAPI(mockCtrl)

	ibmcloudClient.EXPECT().GetVPC(gomock.Any(), validVPC).Return(&vpcv1.VPC{}, nil).AnyTimes()
	ibmcloudClient.EXPECT().GetVPC(gomock.Any(), "not-found").Return(nil, &VPCResourceNotFoundError{})
	ibmcloudClient.EXPECT().GetVPC(gomock.Any(), "internal-error-vpc").Return(nil, fmt.Errorf(""))

	ibmcloudClient.EXPECT().GetSubnet(gomock.Any(), validPublicSubnetUSSouth1ID).Return(&vpcv1.Subnet{Zone: &vpcv1.ZoneReference{Name: &validZoneUSSouth1}}, nil).AnyTimes()
	ibmcloudClient.EXPECT().GetSubnet(gomock.Any(), validPublicSubnetUSSouth2ID).Return(&vpcv1.Subnet{Zone: &vpcv1.ZoneReference{Name: &validZoneUSSouth1}}, nil).AnyTimes()
	ibmcloudClient.EXPECT().GetSubnet(gomock.Any(), validPrivateSubnetUSSouth1ID).Return(&vpcv1.Subnet{Zone: &vpcv1.ZoneReference{Name: &validZoneUSSouth1}}, nil).AnyTimes()
	ibmcloudClient.EXPECT().GetSubnet(gomock.Any(), validPrivateSubnetUSSouth2ID).Return(&vpcv1.Subnet{Zone: &vpcv1.ZoneReference{Name: &validZoneUSSouth1}}, nil).AnyTimes()
	ibmcloudClient.EXPECT().GetSubnet(gomock.Any(), "subnet-invalid-zone").Return(&vpcv1.Subnet{Zone: &vpcv1.ZoneReference{Name: &[]string{"invalid"}[0]}}, nil).AnyTimes()

	ibmcloudClient.EXPECT().GetVSIProfiles(gomock.Any()).Return(validInstanceProfies, nil).AnyTimes()

	ibmcloudClient.EXPECT().GetVPCZonesForRegion(gomock.Any(), validRegion).Return([]string{"us-south-1", "us-south-2", "us-south-3"}, nil).AnyTimes()

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			editedInstallConfig := validInstallConfig()
			for _, edit := range tc.edits {
				edit(editedInstallConfig)
			}

			aggregatedErrors := Validate(ibmcloudClient, editedInstallConfig)
			if tc.errorMsg != "" {
				assert.Regexp(t, tc.errorMsg, aggregatedErrors)
			} else {
				assert.NoError(t, aggregatedErrors)
			}
		})
	}
}

func TestValidatePreExitingPublicDNS(t *testing.T) {
	cases := []struct {
		name     string
		edits    editFunctions
		errorMsg string
	}{
		{
			name:     "no pre-existing DNS records",
			errorMsg: "",
		},
		{
			name:     "pre-existing DNS records",
			errorMsg: `^record api\.valid-cluster-name\.valid\.base\.domain already exists in CIS zone \(valid-zone-id\) and might be in use by another cluster, please remove it to continue$`,
		},
		{
			name:     "cannot get zone ID",
			errorMsg: `^baseDomain: Internal error$`,
		},
		{
			name:     "cannot get DNS records",
			errorMsg: `^baseDomain: Internal error$`,
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	ibmcloudClient := mock.NewMockAPI(mockCtrl)

	dnsRecordName := fmt.Sprintf("api.%s.%s", validClusterName, validBaseDomain)

	// Mocks: no pre-existing DNS records
	ibmcloudClient.EXPECT().GetDNSZoneIDByName(gomock.Any(), validBaseDomain).Return(validDNSZoneID, nil)
	ibmcloudClient.EXPECT().GetDNSRecordsByName(gomock.Any(), validCISInstanceCRN, validDNSZoneID, dnsRecordName).Return(noDNSRecordsResponse, nil)

	// Mocks: pre-existing DNS records
	ibmcloudClient.EXPECT().GetDNSZoneIDByName(gomock.Any(), validBaseDomain).Return(validDNSZoneID, nil)
	ibmcloudClient.EXPECT().GetDNSRecordsByName(gomock.Any(), validCISInstanceCRN, validDNSZoneID, dnsRecordName).Return(existingDNSRecordsResponse, nil)

	// Mocks: cannot get zone ID
	ibmcloudClient.EXPECT().GetDNSZoneIDByName(gomock.Any(), validBaseDomain).Return("", fmt.Errorf(""))

	// Mocks: cannot get DNS records
	ibmcloudClient.EXPECT().GetDNSZoneIDByName(gomock.Any(), validBaseDomain).Return(validDNSZoneID, nil)
	ibmcloudClient.EXPECT().GetDNSRecordsByName(gomock.Any(), validCISInstanceCRN, validDNSZoneID, dnsRecordName).Return(nil, fmt.Errorf(""))

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			validInstallConfig := validInstallConfig()
			meta := &Metadata{
				cisInstanceCRN: validCISInstanceCRN,
			}
			aggregatedErrors := ValidatePreExitingPublicDNS(ibmcloudClient, validInstallConfig, meta)
			if tc.errorMsg != "" {
				assert.Regexp(t, tc.errorMsg, aggregatedErrors)
			} else {
				assert.NoError(t, aggregatedErrors)
			}
		})
	}
}
