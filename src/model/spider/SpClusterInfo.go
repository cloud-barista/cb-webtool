package spider

type SpClusterInfo struct {
	Addons        SpKeyValueList       `json:"Addons"`
	CreatedTime   string               `json:"CreatedTime"`
	IId           SpIID                `json:"IId"`
	Status        string               `json:"Status"`
	Version       string               `json:"Version"`
	KeyValueList  SpKeyValueList       `json:"KeyValueList"`
	Network       SpClusterNetworkInfo `json:"Network"`
	NodeGroupList SpNodeGroupList      `json:"NodeGroupList"`

	ConnectionName string `json:"ConnectionName"` // Spider 미제공으로 추가한 항목.
	ProviderName   string `json:"ProviderName"`   // Spider 미제공으로 추가한 항목.
}

type SpClusterInfoList []SpClusterInfo
