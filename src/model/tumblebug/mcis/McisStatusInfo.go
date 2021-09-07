package mcis

type McisStatusInfo struct {
	ID              string         `json:"id"`
	InstallMonAgent string         `json:"installMonAgent"` // yes, no
	MasterIp        string         `json:"masterIp"`
	MasterSSHPort   string         `json:"masterSSHPort"`
	MasterVmID      string         `json:"masterVmId"`
	Name            string         `json:"name"`
	Status          string         `json:"status"`
	TargetAction    string         `json:"targetAction"`
	TargetStatus    string         `json:"targetStatus"`
	Vm              TbVmStatusInfo `json:"vm"`
}

type McisStatusInfos []McisStatusInfo
