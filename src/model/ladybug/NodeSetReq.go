package ladybug

type NodeSetReq struct {
	Connection string       `json:"connection"`
	Count      int          `json:"count"`
	Spec       string       `json:"spec"`
	RootDisk   RootDiskInfo `json:"rootDisk"`
}
