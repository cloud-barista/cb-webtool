package model

// VM의 상태정보
type VMSimpleInfo struct {
	VmIndex  int    `json:"vmIndex"`
	VmID     string `json:"vmID"`
	VmName   string `json:"vmName"`
	VmStatus string `json:"vmStatus"`

	// Latitude  float64 `json:"latitude"`
	// Longitude float64 `json:"longitude"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
