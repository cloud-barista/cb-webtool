package spider

type ClusterReqInfo struct {
	NameSpace      string      `json:"NameSpace"`
	ConnectionName string      `json:"ConnectionName"`
	ReqInfo        ClusterInfo `json:"ReqInfo"`
}
