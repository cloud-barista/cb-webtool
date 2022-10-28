package spider

type SpVpcInfo struct {
	IId            SpIID        `json:"IId"`
	IPv4_CIDR      string       `json:"IPv4_CIDR"`
	SubnetInfoList SpSubnetList `json:"SubnetInfoList"`
}
