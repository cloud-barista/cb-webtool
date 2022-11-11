package ladybug

// ladybug : cb-mcks 의 model.Cluster
type ClusterInfo struct {
	ClusterConfig string `json:"clusterConfig"`
	CpGroup       string `json:"cpGroup"`
	CpLeader      string `json:"cpLeader"`
	CreatedTime   string `json:"createdTime"`
	Description   string `json:"description"`
	Etcd          string `json:"etcd"` //local, external

	InstallMonAgent string `json:"installMonAgent"`
	K8sVersion      string `json:"k8sVersion"`

	Kind         int           `json:"kind"`
	Label        int           `json:"label"`
	Loadbalancer int           `json:"loadbalancer"` //haproxy, nlb
	Mcis         string        `json:"mcis"`
	Name         string        `json:"name"`
	NameSpace    string        `json:"namespace"`
	NetworkCni   string        `json:"networkCni"` //canal, kilo
	Status       ClusterStatus `json:"status"`
	//UID           string     `json:"uid"`// deprecated
	Nodes []NodeInfo `json:"nodes"`
}
