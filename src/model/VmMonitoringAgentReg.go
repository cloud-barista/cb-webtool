package model

type VmMonitoringAgentReg struct {
	Command  string `json:"command"`
	PublicIp string `json:"ip"`
	McisID   string `json:"mcis_id"`
	SshKey   string `json:"ssh_key"`
	UserName string `json:"user_name"`
	VmID     string `json:"vm_id"`
}
