package spider

type ClusterReqInfo struct {
	NameSpace      string      `json:"Namespace"`
	ConnectionName string      `json:"ConnectionName"`
	ReqInfo        ClusterInfo `json:"ReqInfo"`
}
