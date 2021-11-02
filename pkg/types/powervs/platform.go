package powervs

// Platform stores all the global configuration that all machinesets
// use.
<<<<<<< HEAD
// Note: The subsequent mentions of future-TF support refer to work that is
// undergoing and should be available to test well in time for 4.10 feature-
// freeze. We do not plan to GA with these as required inputs.
type Platform struct {

	// ServiceInstanceID is the ID of the Power IAAS instance created from the IBM Cloud Catalog
	ServiceInstanceID string `json:"serviceInstance"`

	// PowerVSResourceGroup is the resource group for creating Power VS resources.
	PowerVSResourceGroup string `json:"powervsResourceGroup"`

=======
/// used by the installconfig, and filled in by the installconfig/platform/powervs::Platform() func
type Platform struct {
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	// Region specifies the IBM Cloud region where the cluster will be created.
	Region string `json:"region"`

	// Zone specifies the IBM Cloud colo region where the cluster will be created.
	// Required for multi-zone regions.
	Zone string `json:"zone"`

<<<<<<< HEAD
	// UserID is the login for the user's IBM Cloud account.
	UserID string `json:"userID"`

	// APIKey is the API key for the user's IBM Cloud account.
	//
	// +optional
	APIKey string `json:"APIKey,omitempty"`

	// SSHKeyName is the name of an SSH key stored in the Service Instance.
	SSHKeyName string `json:"SSHKeyName,omitempty"`

	// VPC is a VPC inside IBM Cloud. Needed in order to create VPC Load Balancers.
	//
	// @TODO: make this +optional when we have TF support
	VPC string `json:"vpc,omitempty"`
=======
	UserID string `json:"userid"`
	APIKey string `json:"apikey"`
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)

	// Subnets specifies existing subnets (by ID) where cluster
	// resources will be created.  Leave unset to have the installer
	// create subnets in a new VPC on your behalf.
<<<<<<< HEAD
	// @TODO: Rename VPCSubnetID & make into string
	//
	// @TODO: make this +optional when we have TF support
	Subnets []string `json:"subnets,omitempty"`

	// PVSNetworkName specifies an existing network within the Power VS Service Instance.
	// @TODO: make this +optional when we have TF support
	PVSNetworkName string `json:pvsNetworkName,omitempty"`

	// PVSNetworkID is the associated ID for the PVSNetworkName. This is currently required
	// For the machine config.
	// @TODO: Remove when machine config resolves the ID from name.
	// Leave unset to have the installer create the network.
	//
	// @TODO: make this +optional when we have TF support
	PVSNetworkID string `json:"pvsNetworkID,omitempty"`

=======
	// @TODO: how will we handle networking?
	//
	// +optional ?
	Subnets []string `json:"subnets,omitempty"`

>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	// UserTags additional keys and values that the installer will add
	// as tags to all resources that it creates. Resources created by the
	// cluster itself may not include these tags.
	// +optional
	UserTags map[string]string `json:"userTags,omitempty"`

<<<<<<< HEAD
	// ImageName is equivalent to BootStrap/ClusterOSImage.
	// Until the machine provider config in cluster-api-provider-powervs
	// takes an ID instead of a name, we're using this for TF Creation,
	// and the other two for machine-config.
	//
	// @TODO: Remove when provider resolves ID from name
	// @TODO: make this +optional when we have TF support
	ImageName string `json:"imageName,omitempty"`

=======
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	// BootstrapOSImage is a URL to override the default OS image
	// for the bootstrap node. The URL must contain a sha256 hash of the image
	// e.g https://mirror.example.com/images/image.ova.gz?sha256=a07bd...
	//
<<<<<<< HEAD
	// @TODO: make this +optional when we have TF support
	BootstrapOSImage string `json:"bootstrapOSImage,omitempty"`
=======
	// +optional
	BootstrapOSImage string `json:"bootstrapOSImage,omitempty" validate:"omitempty,osimageuri,urlexist"`
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)

	// ClusterOSImage is a URL to override the default OS image
	// for cluster nodes. The URL must contain a sha256 hash of the image
	// e.g https://mirror.example.com/images/powervs.ova.gz?sha256=3b5a8...
	//
<<<<<<< HEAD
	// @TODO: make this +optional when we have TF support
	ClusterOSImage string `json:"clusterOSImage,omitempty"`
=======
	// +optional
	ClusterOSImage string `json:"clusterOSImage,omitempty" validate:"omitempty,osimageuri,urlexist"`
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)

	// DefaultMachinePlatform is the default configuration used when
	// installing on Power VS for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`
}
