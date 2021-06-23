package tumblebug

type McisCommandInfo struct {
	Command  string `json:"command"`
	UserName string `json:"user_name"`
	Ip       string `json:"ip"`
	SshKey   string `json:"ssk_key"`
	McisID   string `json:"mcis_id"`
	VmID     string `json:"vm_id"`
}
