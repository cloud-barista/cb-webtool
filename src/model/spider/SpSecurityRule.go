package spider

type SpSecurityRule struct {
	FromPort   string `json:"FromPort"`
	ToPort     string `json:"ToPort"`
	IPProtocol string `json:"IPProtocol"`
	Direction  string `json:"Direction"`
}

type SpSecurityRuleList []SpSecurityRule
