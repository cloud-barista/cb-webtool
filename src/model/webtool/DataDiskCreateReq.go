package webtool

type DataDiskCreateReq struct {

	//tbmcir.TbDataDiskReq
	Name           string `json:"name"`
	ConnectionName string `json:"connectionName"`
	CspDataDiskId  string `json:"cspDataDiskId"`
	Description    string `json:"description"`
	DiskSize       string `json:"diskSize"`
	DiskType       string `json:"diskType"`

	// Attach VMID
	McisID string `json:"mcisId"`
	//VmID       string `json:"vmId"`
	AttachVmID string `json:"attachVmId"`
}
