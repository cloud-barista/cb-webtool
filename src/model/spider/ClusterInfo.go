package spider

type ClusterInfo struct {
	Name               string          `json:"Name"`
	Version            string          `json:"Version"`
	VPCName            string          `json:"VPCName"`
	SubnetNames        []string        `json:"SubnetNames"`
	SecurityGroupNames []string        `json:"SecurityGroupNames"`
	NodeGroupList      []NodeGroupInfo `json:"NodeGroupList"`
}
