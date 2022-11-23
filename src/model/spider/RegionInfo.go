package spider

type RegionInfo struct {
	RegionName       string         `json:"RegionName"`
	ProviderName     string         `json:"ProviderName"`
	KeyValueInfoList SpKeyValueList `json:"KeyValueInfoList"`
}

type RegionInfos []RegionInfo
