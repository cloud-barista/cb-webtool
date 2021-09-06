package common

type TbRegion struct {
	ProviderName     string       `json:"providerName"`
	RegionName       string       `regionName:"description"`
	KeyValueInfoList []TbKeyValue `json:"resourcesOnTumblebug"`
}

type TbRegions []TbRegion
