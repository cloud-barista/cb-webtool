package mcir

type TbCustomImageReq struct {
	ConnectionName   string `json:"connectionName"`
	CspCustomImageId string `json:"cspCustomImageId"`

	Description string `json:"description"`
	Name        string `json:"name"`

	SourceVmID string `json:"sourceVmId"`
}
