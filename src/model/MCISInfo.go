package model

type MCISInfo struct {
	// ID     string `json:"id"`
	// Name   string `json:"name"`
	// Status string `json:"status"`
	// VMNum  string `json:"vm_num"`

	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`

	InstallMonAgent string `json:"installMonAgent"`
	Label           string `json:"label"`

	PlacementAlgo string `json:"placement_algo"`

	TargetAction string `json:"targetAction"`
	TargetStatus string `json:"targetStatus"`

	VMs []VMInfo `json:"vm"`
}

// MCIS의 일부정보만 추려서
type MCISSimpleInfo struct {
	// ID     string `json:"id"`
	// Name   string `json:"name"`
	// Status string `json:"status"`
	// VMNum  string `json:"vm_num"`

	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`

	ConnectionCount   int `json:"connectionCount"`
	VmCount           int `json:"vmCount"`
	VmRunningCount    int `json:"vnRunningCount"`
	VmStoppedCount    int `json:"vmStopped"`
	VmTerminatedCount int `json:"vmTerminated"`
	// mcis.ID, mcis.status, mcis.name, mcis.description
	// csp : 해당 mcis의 connection cnt
	// vm_cnt : 해당 mcis의 vm cnt
	// vm_run_cnt, vm_stop_cnt
}
type MCISSimpleInfos []MCISSimpleInfo
