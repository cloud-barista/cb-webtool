package tumblebug

type SecurityGroupInfo struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ConnectionName string `json:"connectionName"`
	Description    string `json:"description"`
	VNetID         string `json:"vNetID"`

	CspSecurityGroupId   string `json:"cspSecurityGroupId"`
	CspSecurityGroupName string `json:"cspSecurityGroupName"`

	FirewallRules []FirewallRule `json:"firewallRules"`

	KeyValueInfos []KeyValueInfo `json:"keyValueList"`
}

type SecurityGroupInfos []SecurityGroupInfo
