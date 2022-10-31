package mcir

type RestGetAvailableDataDisksResponse struct {
	//DataDisk []string `json:"dataDisk"`
	DataDisk []TbDataDiskInfo `json:"dataDisk"`
}
