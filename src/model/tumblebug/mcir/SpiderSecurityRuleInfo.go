package mcir

type SpiderSecurityRuleInfo struct {
	Cidr        string `json:"cidr"`
	Description string `json:"description"`
	fromPort    string `json:"fromPort"`
	toPort      string `json:"toPort"`
	ipProtocol  string `json:"ipProtocol"`
}
