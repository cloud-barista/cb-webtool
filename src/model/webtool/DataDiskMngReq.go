package webtool

import (
	tbmcir "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcir"
)

type DataDiskMngReq struct {
	// create disk list
	CreateDataDiskList []DataDiskCreateReq `json:"createDataDiskList"`

	// attach list
	AttachDataDiskList []tbmcir.TbAttachDetachDataDiskReq `json:"attachDataDiskList"`

	// detach list
	DetachDataDiskList []tbmcir.TbAttachDetachDataDiskReq `json:"dettachDataDiskList"`

	// del disk list
	DeleteDataDiskList []string `json:"dataDiskList"`

	// update :TbDataDiskUpsizeReq
}
