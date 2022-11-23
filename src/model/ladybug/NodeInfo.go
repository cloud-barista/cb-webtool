package ladybug

// ladybug
type NodeInfo struct {
	CreatedTime string `json:"createdTime"`
	Credential  string `json:"credential"`
	Csp         string `json:"csp"`
	CspLabel    string `json:"cspLable"`
	ZoneLabel   string `json:"zoneLable"`
	Kind        string `json:"kind"` // Node 냐 cluster냐
	Name        string `json:"name"`
	PublicIp    string `json:"publicIp"`
	Role        string `json:"role"` // Control-plane냐, Worker냐
	Spec        string `json:"spec"`
	//UID        string `json:"uid"`
}
type Nodes []NodeInfo
