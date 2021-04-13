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

type MCISInfos []MCISInfo

// MCIS의 일부정보만 추려서
type MCISSimpleInfo struct {
	// ID     string `json:"id"`
	// Name   string `json:"name"`
	// Status string `json:"status"`
	// VMNum  string `json:"vm_num"`

	// mcis.ID, mcis.status, mcis.name, mcis.description
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	McisStatus  string `json:"mcisStatus"`
	Description string `json:"description"`

	ConnectionCount int `json:"connectionCount"`

	// vm_cnt : 해당 mcis의 vm cnt
	// vm_run_cnt, vm_stop_cnt
	VmCount          int            `json:"vmCount"`
	VmStatusNames    string         `json:"vmStatusNames"`
	VmStatusList     []VMStatus     `json:"vmStatusList"`
	VmStatusCountMap map[string]int `json:"vmStatusCountMap"`
	// VmRunningCount    int `json:"vnRunningCount"`
	// VmStoppedCount    int `json:"vmStopped"`
	// VmTerminatedCount int `json:"vmTerminated"`

	// csp : 해당 mcis의 connection cnt
	ConnectionConfigProviderMap   map[string]int `json:"connectionConfigProviderMap"`
	ConnectionConfigProviderNames string         `json:"connectionConfigProviderNames"` // 해당 MCIS 등록된 connection의 provider 목록
	// ConnectionConfigProviderNames []string       `json:"connectionConfigProviderNames"` // 해당 MCIS 등록된 connection의 provider 목록
	ConnectionConfigProviderCount int `json:"connectionConfigProviderCount"`
}
type MCISSimpleInfos []MCISSimpleInfo
