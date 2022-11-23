package mcir

type TbDataDiskUpsizeReq struct {
	Description string `json:"description"`
	DiskSize    string `json:"diskSize"`
}
