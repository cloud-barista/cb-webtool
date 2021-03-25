package model

type SecurityGroupRegInfo struct {
	Name           string `json:"name"`
	ConnectionName string `json:"connectionName"`
	Description    string `json:"description"`
	VNetID         string `json:"vNetID"`

	FirewallRules FirewallRules `json:"firewallRules"`
}

type SecurityGroupRegInfos []SecurityGroupRegInfo

type FirewallRules struct {
	Direction  string `json:"direction"`
	FromPort   string `json:"fromPort"`
	IpProtocol string `json:"ipProtocol"`
	ToPort     string `json:"toPort"`
}
