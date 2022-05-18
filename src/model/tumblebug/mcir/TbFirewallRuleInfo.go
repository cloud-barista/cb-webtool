package mcir

type TbFirewallRuleInfo struct {
	Cidr       string `json:"cidr"`
	Direction  string `json:"direction"`//require
	FromPort   string `json:"fromPort"`//require
	ToPort     string `json:"toPort"`//require
	IpProtocol string `json:"ipprotocol"`//require
}