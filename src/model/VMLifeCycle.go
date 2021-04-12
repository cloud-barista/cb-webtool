package model

// Life Cycle command 전송용
type VMLifeCycle struct {
	NameSpaceID        string `json:"nameSpaceID"`
	McisID             string `json:"mcisID"`
	VmID               string `json:"vmID"`
	LifeCycleOperation string `json:"lifeCycleOperation"`
}
