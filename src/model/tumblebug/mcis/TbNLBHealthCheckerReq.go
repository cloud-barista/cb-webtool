package mcis

type TbNLBHealthCheckerReq struct {
	// Protocol string `json:"protocol"`
	// Port     string `json:"port"`

	Interval  string `json:"interval"`
	Threshold string `json:"threshold"`
	Timeout   string `json:"timeout"`
}
