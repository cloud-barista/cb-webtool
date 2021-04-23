package model

type McisRequest struct {
	VmSpec           []string `form:"vmspec"`
	NameSpace        string   `form:"namespace"`
	McisName         string   `form:"mcis_name"`
	VmName           []string `form:"vmName"`
	Provider         []string `form:"provider"`
	SecurityGroupIds []string `form:"sg"`
}
