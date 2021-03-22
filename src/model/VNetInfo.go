package model

type VNetInfo struct {
	cidrBlock      string       `json:"cidrBlock"`
	connectionName string       `json:"connectionName"`
	description    string       `json:"description"`
	name           string       `json:"name"`
	SubnetInfos    []SubnetInfo `json:"subnetInfoList"`
}

type VNetInfos []VNetInfo
