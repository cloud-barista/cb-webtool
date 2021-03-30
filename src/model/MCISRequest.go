package model

type MCISRequest struct {
	VMSpec           []string `form:"vmspec"`
	NameSpace        string   `form:"namespace"`
	McisName         string   `form:"mcis_name"`
	VMName           []string `form:"vmName"`
	Provider         []string `form:"provider"`
	SecurityGroupIds []string `form:"sg"`
}