package mcis

type TbMcisReq struct {
	Description     string `json:"description"`
	InstallMonAgent string `json:"installMonAgent"`
	Label           string `json:"label"`
	Name            string `json:"name"`
	PlacementAlgo   string `json:"placementAlgo"`

	Vm TbVmInfo `json:"vm"`
}
