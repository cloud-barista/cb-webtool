package model

// VM의 상태정보
type VMStatus struct {
	VmIndex  int    `json:"vmIndex"`
	VmID     string `json:"vmID"`
	VmName   string `json:"vmName"`
	VmStatus string `json:"vmStatus"`
}
