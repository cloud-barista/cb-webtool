package mcis

type TbInspectResourcesResponse struct {
	resourcesOnCsp []string `json:"resourcesOnCsp"`
	// resourcesOnSpider    ResourceOnCspOrSpider `json:"resourcesOnSpider"`
	// resourcesOnTumblebug ResourceOnTumblebug   `json:"resourcesOnTumblebug"`
}
