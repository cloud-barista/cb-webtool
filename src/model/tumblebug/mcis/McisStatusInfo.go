package mcis

type McisStatusInfo struct {
	id              string         `json:"maxResultNum"`
	installMonAgent string         `json:"installMonAgent"` // yes, no
	masterIp        string         `json:"masterIp"`
	masterSSHPort   string         `json:"masterSSHPort"`
	masterVmId      string         `json:"masterVmId"`
	name            string         `json:"name"`
	status          string         `json:"status"`
	targetAction    string         `json:"targetAction"`
	targetStatus    string         `json:"targetStatus"`
	vm              TbVmStatusInfo `json:"vm"`
}

type McisStatusInfos []McisStatusInfo
