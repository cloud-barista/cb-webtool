package mcis

type TbNLBAddRemoveVMReq struct {
	TargetGroup []TbNLBTargetGroup `json:"targetGroup"`
}
