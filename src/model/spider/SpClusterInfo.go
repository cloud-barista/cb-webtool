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
	AccessInfo    SpClusterAccessInfo  `json:"AccessInfo"`

	ConnectionName string `json:"ConnectionName"` // Spider 미제공으로 추가한 항목.
	ProviderName   string `json:"ProviderName"`   // Spider 미제공으로 추가한 항목.
}

type SpClusterInfoList []SpClusterInfo

// Connection이 다른 namespace내 모든 Cluster의 List
type SpAllClusterInfoList struct {
	Connection string `json:"Connection"`
	Provider   string `json:"Provider"`

	ClusterList SpClusterInfoList `json:"ClusterList"`
	//ClusterList []SpClusterInfo `json:"ClusterList"`
}

type SpTotalClusterInfoList struct {
	AllClusterList []SpAllClusterInfoList `json:"AllClusterList"`
}

type RespClusterInfo struct {
	ClusterInfo SpClusterInfo `json:"ClusterInfo"`
}
