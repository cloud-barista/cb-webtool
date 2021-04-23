package model

type FirewallRule struct {
	Direction  string `json:"direction"`
	FromPort   string `json:"fromPort"`
	IpProtocol string `json:"ipProtocol"`
	ToPort     string `json:"toPort"`
}
