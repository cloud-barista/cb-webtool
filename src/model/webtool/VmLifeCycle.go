package webtool

// Life Cycle command 전송용
type VmLifeCycle struct {
	NameSpaceID   string `json:"nameSpaceID"`
	McisID        string `json:"mcisID"`
	VmID          string `json:"vmID"`
	LifeCycleType string `json:"lifeCycleType"` // reboot, create, suspend, resume, terminate, delete  : Const.VM_LIFECYCLE_xxx
}
