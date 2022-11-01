package mcis

type TbVmReq struct {
	ConnectionName string `json:"connectionName"`
	Description    string `json:"description"`
	IdByCsp        string `json:"idByCsp"`
	ImageID        string `json:"imageId"`
	ImageType      string `json:"imageType"`
	Label          string `json:"label"`
	Name           string `json:"name"`

	RootDiskSize string `json:"rootDiskSize"`
	RootDiskType string `json:"rootDiskType"`

	SecurityGroupIDs []string `json:"securityGroupIds"`

	SpecID         string `json:"specId"`
	SshKeyID       string `json:"sshKeyId"`
	SubnetID       string `json:"subnetId"`
	VNetID         string `json:"vNetId"`
	SubGroupSize   string `json:"subGroupSize"`
	VmUserAccount  string `json:"vmUserAccount"`
	VmUserPassword string `json:"vmUserPassword"`

	DataDiskIds []string `json:"dataDiskIds"`
}
