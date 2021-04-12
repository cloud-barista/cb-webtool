package model

// Monitoring 수신 Data  return용
type VMMonitoringInfo struct {
	Name       string              `json:"name"`
	Tags       VMMonitoringTag     `json:"tags"`
	ValuesList []VMMonitoringValue `json:"values"`
}

type VMMonitoringTag struct {
	VmID string `json:"vmId"`
}

type VMMonitoringValue struct {
	CpuGuest       float64 `json:"cpu_guest"`
	CpuGuestNice   float64 `json:"cpu_guest_nice"`
	CpuHintr       float64 `json:"cpu_hintr"`
	CpuIdle        float64  `json:"cpu_idle"`
	CpuIowait      float64 `json:"cpu_iowait"`
	CpuNice        float64 `json:"cpu_nice"`
	CpuSintr       float64 `json:"cpu_sintr"`
	CpuSteal       float64 `json:"cpu_steal"`
	CpuSystem      float64  `json:"cpu_system"`
	CpuUser        float64 `json:"cpu_user"`
	CpuUtilization float64 `json:"cpu_utilization"`
	Time           string `json:"time"`

}
