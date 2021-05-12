package ladybug

// ladybug
type ClusterInfo struct {
	ClusterConfig string     `json:"clusterConfig"`
	Kind          int        `json:"kind"`
	Mcis          string     `json:"mcis"`
	Name          string     `json:"name"`
	NameSpace     string     `json:"namespace"`
	Status        string     `json:"status"`
	UID           string     `json:"uid"`
	Nodes         []NodeInfo `json:"nodes"`
}
