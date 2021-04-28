package ladybug

// ladybug
type NodeInfo struct {
	Credential string `json:"credential"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	PublicIp   string `json:"publicIp"`
	Role       string `json:"role"`
	Spec       string `json:"spec"`
	UID        string `json:"uid"`
}
type Nodes []NodeInfo
