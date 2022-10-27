package spider

type SpSubnetInfo struct {
	IId          SpIID            `json:"IId"`
	IPv4_CIDR    string           `json:"IPv4_CIDR"`
	KeyValueList []SpKeyValueInfo `json:"KeyValueList"`
}
type SpSubnetList []SpSubnetInfo
