package spider

type SpNodeGroupInfo struct {
	IId        SpIID `json:"IId"`
	ImageIID   SpIID `json:"ImageIID"`
	KeyPairIID SpIID `json:"KeyPairIID"`

	DesiredNodeSize int       `json:"DesiredNodeSize"`
	MaxNodeSize     int       `json:"MaxNodeSize"`
	MinNodeSize     int       `json:"MinNodeSize"`
	Nodes           SpIIDList `json:"Nodes"`

	OnAutoScaling bool   `json:"OnAutoScaling"`
	RootDiskSize  int    `json:"RootDiskSize"`
	RootDiskType  string `json:"RootDiskType"`
	Status        string `json:"Status"`
	VMSpecName    string `json:"VMSpecName"`
}

type SpNodeGroupList []SpNodeGroupInfo
