package spider

type SpSecurityGroupInfo struct {
	IId           SpIID              `json:"IId"`
	VpcIID        SpIID              `json:"VpcIId"`
	Direction     string             `json:"Direction"`
	SecurityRules SpSecurityRuleList `json:"SecurityRules"`
	KeyValueList  []SpKeyValueInfo   `json:"KeyValueList"`
}

type SpSecurityGroupList []SpSecurityGroupInfo
