package spider

type NodeGroupReqInfo struct {
	NameSpace      string        `json:"NameSpace"`
	ConnectionName string        `json:"ConnectionName"`
	ReqInfo        NodeGroupInfo `json:"ReqInfo"`
}
