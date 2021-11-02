package powervs

<<<<<<< HEAD
// MachinePool stores the configuration for a machine pool installed on IBM Power VS.
=======
// MachinePool stores the configuration for a machine pool installed on Ibm Power VS.
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
type MachinePool struct {
	// ServiceInstance is Service Instance to install into.
	//
	ServiceInstance string `json:"serviceinstance"`

	// Name is the name of the instance
	//
	Name string `json:"name"`

<<<<<<< HEAD
	// KeyPairName is the name of an SSH key pair stored in the Power VS
	// Service Instance
	KeyPairName string `json:"keypairname"`

=======
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	// VolumeIDs is the list of volumes attached to the instance.
	//
	VolumeIDs []string `json:"volumeIDs"`

	// Memory defines the memory in GB for the instance.
	//
	Memory int `json:"memory"`

	// Processors defines the processing units for the instance.
	// @TODO:
	Processors float32 `json:"processors"`

	// ProcType defines the processor sharing model for the instance.
	//
	// +optional
	ProcType string `json:"procType"`

	// ImageID defines the ImageID for the instance.
	//
	// +optional (does this mean user-optional, or completely?)
	ImageID string `json:"imageID"`

	// NetworkIDs defines the network IDs of the instance.
	//
	// +optional
	NetworkIDs []string `json:"networkIDs"`

<<<<<<< HEAD
=======
	// KeyPairName defines the keyPairName name for instance.
	//
	// +optional
	KeyPairName string `json:"keyPairName"`

>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
	// SysType defines the system type for instance.
	//
	// +optional
	SysType string `json:sysType"`
}

// Set stores values from required into a
func (a *MachinePool) Set(required *MachinePool) {
	if required == nil || a == nil {
		return
	}
	if required.ImageID != "" {
		a.ImageID = required.ImageID
	}
	if required.ServiceInstance != "" {
		a.ServiceInstance = required.ServiceInstance
	}
<<<<<<< HEAD
	if len(required.NetworkIDs) > 0 {
		a.NetworkIDs = required.NetworkIDs
	}
	if required.ServiceInstance != "" {
		a.ServiceInstance = required.ServiceInstance
	}
	if required.KeyPairName != "" {
		a.KeyPairName = required.KeyPairName
	}
=======
	if required.KeyPairName != "" {
		a.KeyPairName = required.KeyPairName
	}
	if len(required.NetworkIDs) > 0 {
		a.NetworkIDs = required.NetworkIDs
	}
	// Sets values taken from passed MachinePool.
	/*
		if len(required.Zones) > 0 {
			a.Zones = required.Zones
		}

		if required.InstanceType != "" {
			a.InstanceType = required.InstanceType
		}

		if required.OSDisk.DiskSizeGB > 0 {
			a.OSDisk.DiskSizeGB = required.OSDisk.DiskSizeGB
		}

		if required.OSDisk.DiskType != "" {
			a.OSDisk.DiskType = required.OSDisk.DiskType
		}

		if required.EncryptionKey != nil {
			if a.EncryptionKey == nil {
				a.EncryptionKey = &EncryptionKeyReference{}
			}
			a.EncryptionKey.Set(required.EncryptionKey)
		}*/
>>>>>>> ce5d7615b (Squashing Power VS IPI commits)
}
