package mcis

import tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"

type TbNLBTargetGroup struct {
	CspID string `json:"cspID"`

	Protocol string `json:"protocol"`
	Port     string `json:"port"`

	SubGroupId string   `json:"subGroupId"`
	Vms        []string `json:"vms"`

	KeyValueList []tbcommon.TbKeyValue `json:"keyValueList"`
}
