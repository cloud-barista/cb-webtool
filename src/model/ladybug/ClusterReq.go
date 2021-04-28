package ladybug

// ladybug
type ClusterReq struct {
	ControlPlaneNodeCount int    `json:"controlPlaneNodeCount"`
	ControlPlaneNodeSpec  string `json:"controlPlaneNodeSpec"`
	Name                  string `json:"name"`
	WorkerNodeCount       int    `json:"workerNodeCount"`
	WorkerNodeSpec        string `json:"workerNodeSpec"`
}
