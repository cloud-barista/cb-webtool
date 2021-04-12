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
	// CpuGuest       string `json:"cpu_guest"`
	// CpuGuestNice   string `json:"cpu_guest_nice"`
	// CpuHintr       string `json:"cpu_hintr"`
	// CpuIdle        string `json:"cpu_idle"`
	// CpuIowait      string `json:"cpu_iowait"`
	// CpuNice        string `json:"cpu_nice"`
	// CpuSintr       string `json:"cpu_sintr"`
	// CpuSteal       string `json:"cpu_steal"`
	// CpuSystem      string `json:"cpu_system"`
	// CpuUser        string `json:"cpu_user"`
	// CpuUtilization string `json:"cpu_utilization"`
	// Time           string `json:"time"`

	CpuGuest       string `json:"cpuGuest"`
	CpuGuestNice   string `json:"cpu_guest_nice"`
	CpuHintr       string `json:"cpuHintr"`
	CpuIdle        string `json:"cpuIdle"`
	CpuIowait      string `json:"cpu_iowait"`
	CpuNice        string `json:"cpu_nice"`
	CpuSintr       string `json:"cpuSintr"`
	CpuSteal       string `json:"cpu_steal"`
	CpuSystem      string `json:"cpuSystem"`
	CpuUser        string `json:"cpu_user"`
	CpuUtilization string `json:"cpu_utilization"`
	Time           string `json:"time"`
}
