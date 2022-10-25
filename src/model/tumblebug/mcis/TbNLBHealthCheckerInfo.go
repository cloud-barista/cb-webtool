package mcis

import tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"

type TbNLBHealthCheckerInfo struct {
	CspID string `json:"cspID"`

	Protocol string `json:"protocol"`
	Port     string `json:"port"`

	Interval  int `json:"interval"`
	Threshold int `json:"threshold"`
	Timeout   int `json:"timeout"`

	KeyValueList []tbcommon.TbKeyValue `json:"keyValueList"`
}
