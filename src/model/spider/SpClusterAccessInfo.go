package spider

type SpClusterAccessInfo struct {
	Endpoint   string `json:"Endpoint"` //ex)https://1.2.3.4:56
	Kubeconfig string `json:"Kubeconfig"`
}
