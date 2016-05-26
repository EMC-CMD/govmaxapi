package model

type EditStorageGroupParam struct {
	EditStorageGroupActionParam EditStorageGroupActionParam `json:"editStorageGroupActionParam"`
}

type EditStorageGroupActionParam struct {
	ExpandStorageGroupParam ExpandStorageGroupParam `json:"expandStorageGroupParam,omitempty"`
	RemoveVolumeParam       RemoveVolumeParam       `json:"removeVolumeParam,omitempty"`
}

type ExpandStorageGroupParam struct {
	NumOfVols        int             `json:"num_of_vols"`
	VolumeAttribute  VolumeAttribute `json:"volumeAttribute"`
	CreateNewVolumes bool            `json:"create_new_volumes"`
	Emulation        string          `json:"emulation"`
}

type VolumeAttribute struct {
	CapacityUnit string `json:"capacityUnit"`
	VolumeSize   string `json:"volume_size"`
}

type EditStorageGroupResponse struct {
	StorageGroupID    string  `json:"storageGroupId"`
	Slo               string  `json:"slo"`
	Srp               string  `json:"srp"`
	Workload          string  `json:"workload"`
	SloCompliance     string  `json:"slo_compliance"`
	NumOfVols         int     `json:"num_of_vols"`
	NumOfChildSgs     int     `json:"num_of_child_sgs"`
	NumOfParentSgs    int     `json:"num_of_parent_sgs"`
	NumOfMaskingViews int     `json:"num_of_masking_views"`
	NumOfSnapshots    int     `json:"num_of_snapshots"`
	CapGb             float64 `json:"cap_gb"`
	DeviceEmulation   string  `json:"device_emulation"`
	Type              string  `json:"type"`
	Unprotected       bool    `json:"unprotected"`
}

type RemoveVolumeParam struct {
	VolumeID []string `json:"volumeId"`
}
