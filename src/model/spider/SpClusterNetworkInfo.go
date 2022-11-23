package spider

type SpClusterNetworkInfo struct {
	VpcIID            SpIID     `json:"VpcIID"`
	SubnetIIDs        SpIIDList `json:"SubnetIIDs"`
	SecurityGroupIIDs SpIIDList `json:"SecurityGroupIIDs"`

	KeyValueList SpKeyValueList `json:"KeyValueList"`
}
