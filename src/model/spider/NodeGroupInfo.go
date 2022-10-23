package spider

type NodeGroupInfo struct {
	Name            string `json:"Name"`
	Version         string `json:"Version"`
	ImageName       string `json:"ImageName"`
	VMSpecName      string `json:"VMSpecName"`
	KeyPairName     string `json:"KeyPairName"`
	OnAutoScaling   string `json:"OnAutoScaling"`
	DesiredNodeSize string `json:"DesiredNodeSize"`
	MinNodeSize     string `json:"MinNodeSize"`
	MaxNodeSize     string `json:"MaxNodeSize"`
}
