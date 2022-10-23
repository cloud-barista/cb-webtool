package spider

type ClusterInfo struct {
	NameSpace      string         `json:"NameSpace"`
	ConnectionName string         `json:"ConnectionName"`
	ReqInfo        ClusterReqInfo `json:"ReqInfo"`
}
