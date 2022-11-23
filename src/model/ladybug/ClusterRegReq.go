package ladybug

// ladybug : ClusterReq
type ClusterRegReq struct {
	// ControlPlaneNodeCount int    `json:"controlPlaneNodeCount"`
	// ControlPlaneNodeSpec  string `json:"controlPlaneNodeSpec"`
	// Name                  string `json:"name"`
	// WorkerNodeCount       int    `json:"workerNodeCount"`
	// WorkerNodeSpec        string `json:"workerNodeSpec"`

	Name         string           `json:"name"`
	Label        string           `json:"label"`
	Description  string           `json:"description"`
	Config       ClusterConfigReq `json:"config"`
	ControlPlane []NodeSetReq     `json:"controlPlane"`
	Worker       []NodeSetReq     `json:"worker"`
}
