package mcir

type TbDataDiskReq struct {
	ConnectionName string `json:"connectionName"`
	CspDataDiskId  string `json:"cspDataDiskId"`
	Description    string `json:"description"`
	DiskSize       string `json:"diskSize"`
	DiskType       string `json:"diskType"`

	Name string `json:"name"`
}