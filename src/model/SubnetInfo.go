package model

type SubnetInfo struct {
	IID           IIDInfo        `json:"IId"`
	Ipv4_CIDR     string         `json:"ipv4_CIDR"`
	KeyValueInfos []KeyValueInfo `json:"keyValueList"`
}

type SubnetInfos []SubnetInfo
