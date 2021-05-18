package ladybug

// ClusterInfo의 간단버전.
type ClusterSimpleInfo struct {
	ClusterConfig string     `json:"clusterConfig"`
	CpLeader      int        `json:"cpLeader"`
	Kind          int        `json:"kind"`
	Mcis          string     `json:"mcis"`
	Name          string     `json:"name"`
	NameSpace     string     `json:"namespace"`
	NetworkCni    string     `json:"networkCni"`
	Status        string     `json:"status"`
	McisStatus 	  string     `json:"mcisStatus"`
	UID           string     `json:"uid"`
	Nodes         []NodeSimpleInfo `json:"nodes"`
	TotalNodeCount int `json:"totalNodeCount"`	
}
