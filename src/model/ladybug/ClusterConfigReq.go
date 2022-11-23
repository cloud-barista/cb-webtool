package ladybug

type ClusterConfigReq struct {
	InstallMonAgent string                     `json:"installMonAgent"`
	Kubernetes      ClusterConfigKubernetesReq `json:"kubernetes"`
}

type ClusterConfigKubernetesReq struct {
	Etcd         string `json:"etcd"`
	Loadbalancer string `json:"loadbalancer"`

	NetworkCni       string `json:"networkCni"`
	PodCidr          string `json:"podCidr"`
	ServiceCidr      string `json:"serviceCidr"`
	ServiceDnsDomain string `json:"serviceDnsDomain"`

	Storageclass ClusterStorageClassNfsReq `json:"storageclass"`
	Version      string                    `json:"version"`
}

type ClusterStorageClassNfsReq struct {
	Path   string `json:"path"`
	Server string `json:"server"`
}
