package mcis

import tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"

type TbNLBHealthCheckerInfo struct {
	CspID string `json:"cspID"`

	Protocol string `json:"protocol"`
	Port     string `json:"port"`

	Interval  string `json:"interval"`
	Threshold string `json:"threshold"`
	Timeout   string `json:"timeout"`

	KeyValueList []tbcommon.TbKeyValue `json:"keyValueList"`
}
