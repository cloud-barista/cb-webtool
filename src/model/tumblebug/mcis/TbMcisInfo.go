package mcis

type TbMcisInfo struct {
	ID              string `json:"id"`
	Description     string `json:"description"`
	InstallMonAgent string `json:"installMonAgent"`
	Label           string `json:"label"`
	Name            string `json:"name"`
	PlacementAlgo   string `json:"placementAlgo"`
	Status          string `json:"status"`
	TargetAction    string `json:"targetAction"`
	TargetStatus    string `json:"targetStatus"`

	Vm []TbVmInfo `json:"vm"`
}
