package webtool

type DataDiskMngReq struct {
	// create disk list
	CreateDataDiskList []DataDiskCreateReq `json:"createDataDiskList"`

	// attach list
	AttachDataDiskList []string `json:"attachDataDiskList"`

	// detach list
	DetachDataDiskList []string `json:"dettachDataDiskList"`

	// del disk list
	DeleteDataDiskList []string `json:"dataDiskList"`

	// update :TbDataDiskUpsizeReq
}
